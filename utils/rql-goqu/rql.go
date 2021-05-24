package rqlgq

import (
	"container/list"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

//go:generate easyjson -omit_empty -disallow_unknown_fields -snake_case rql.go

// Query is the decoded result of the user input.
//easyjson:json
type Query struct {
	// Limit must be > 0 and <= to `LimitMaxValue`.
	Limit uint `json:"limit,omitempty"`
	// Offset must be >= 0.
	Offset uint `json:"offset,omitempty"`
	// Select contains the list of expressions define the value for the `SELECT` clause.
	// For example:
	//
	//	params, err := p.Parse([]byte(`{
	//		"select": ["name", "age"]
	//	}`))
	//
	Select []string `json:"select,omitempty"`
	// Sort contains list of expressions define the value for the `ORDER BY` clause.
	// In order to return the rows in descending order you can prefix your field with `-`.
	// For example:
	//
	//	params, err := p.Parse([]byte(`{
	//		"sort": ["name", "-age", "+redundant"]
	//	}`))
	//
	Sort []string `json:"sort,omitempty"`
	// Filter is the query object for building the value for the `WHERE` clause.
	// The full documentation of the supported operators is writtern in the README.
	// An example for filter object:
	//
	//	params, err := p.Parse([]byte(`{
	//		"filter": {
	//			"account": { "$like": "%github%" },
	//			"$or": [
	//				{ "city": "TLV" },
	//				{ "city": "NYC" }
	//			]
	//		}
	//	}`))
	//
	Filter map[string]interface{} `json:"filter,omitempty"`
}

// Params is the parser output after calling to `Parse`. You should pass its
// field values to goqu query tool. For example:
//	params, err := p.Parse(bytez)
//	if err != nil {
//		log.Fatal("Error", err)
//	}
//	sql, args, _  := goqu.From("users").
//		Select(params.Select).
//		Where(params.FilterExp).
//		Limit(params.Limit).
//		Offset(params.Offset).
//		Order(params.Sort...).
//		OrderPrepend(goqu.I("a_column").Asc()).
//
type Params struct {
	// Limit represents the number of rows returned by the SELECT statement.
	Limit uint
	// Offset specifies the offset of the first row to return. Useful for pagination.
	Offset uint
	// Select contains the expression for the `SELECT` clause defined in the Query.
	Select exp.ColumnListExpression
	// Sort used as a parameter for the `ORDER BY` clause. For example, "age desc, name".
	Sort []exp.OrderedExpression
	// FilterExp and FilterArgs come together and used as a parameters for the `WHERE` clause.
	//
	// examples:
	// 	1. Exp: "name = ?"
	//	   Args: "a8m"
	//
	//	2. Exp: "name = ? AND age >= ?"
	// 	   Args: "a8m", 22
	FilterExp  exp.ExpressionList
	FilterArgs []interface{}
}

// ParseError is type of error returned when there is a parsing problem.
type ParseError struct {
	msg string
}

func (p ParseError) Error() string {
	return p.msg
}

// field is a configuration of a struct field.
type field struct {
	// Alias of the column.
	Alias string
	// Column name of the column.
	Column string
	// Has a "sort" option in the tag.
	Sortable bool
	// Has a "filter" option in the tag.
	Filterable bool
	// All supported operators for this field.
	FilterOps map[string]bool
	// Validation for the type. for example, unit8 greater than or equal to 0.
	ValidateFn func(interface{}) error
	// ConvertFn converts the given value to the type value.
	CovertFn func(interface{}) interface{}
}

// A Parser parses various types. The result from the Parse method is a Param object.
// It is safe for concurrent use by multiple goroutines except for configuration changes.
type Parser struct {
	Config
	fields map[string]*field
}

// NewParser creates a new Parser. it fails if the configuration is invalid.
func NewParser(c Config) (*Parser, error) {
	if err := c.defaults(); err != nil {
		return nil, err
	}
	p := &Parser{
		Config: c,
		fields: make(map[string]*field),
	}
	if err := p.init(); err != nil {
		return nil, err
	}
	return p, nil
}

// MustNewParser is like NewParser but panics if the configuration is invalid.
// It simplifies safe initialization of global variables holding a resource parser.
func MustNewParser(c Config) *Parser {
	p, err := NewParser(c)
	if err != nil {
		panic(err)
	}
	return p
}

// Parse parses the given buffer into a Param object. It returns an error
// if the JSON is invalid, or its values don't follow the schema of rql.
func (p *Parser) Parse(b []byte) (pr *Params, err error) {
	q := &Query{}
	if err := q.UnmarshalJSON(b); err != nil {
		return nil, &ParseError{"decoding buffer to *Query: " + err.Error()}
	}
	return p.ParseQuery(q)
}

// ParseQuery parses the given struct into a Param object. It returns an error
// if one of the query values don't follow the schema of rql.
func (p *Parser) ParseQuery(q *Query) (pr *Params, err error) {
	defer func() {
		if e := recover(); e != nil {
			perr, ok := e.(*ParseError)
			if !ok {
				panic(e)
			}
			err = perr
			pr = nil
		}
	}()
	pr = &Params{}
	if p.UseDefault {
		pr.Limit = p.DefaultLimit
	}

	pr.Offset = q.Offset
	if q.Limit != 0 {
		expect(q.Limit <= p.LimitMaxValue, "limit must be greater than 0 and less than or equal to %d", p.LimitMaxValue)
		pr.Limit = q.Limit
	}
	ps := p.newParseState()
	pr.FilterExp = ps.and(q.Filter)
	pr.FilterArgs = ps.values
	pr.Sort = p.sort(q.Sort)
	if len(pr.Sort) == 0 && len(p.DefaultSort) > 0 {
		pr.Sort = p.sort(p.DefaultSort)
	}

	s := make([]interface{}, len(q.Select))
	for i := range s {
		s[i] = q.Select[i]
	}
	pr.Select = exp.NewColumnListExpression(s...)

	parseStatePool.Put(ps)
	return
}

// Column is the default function that converts field name into a database column.
// It used to convert the struct fields into their database names. For example:
//
//	Username => username
//	FullName => full_name
//	HTTPCode => http_code
//
func Column(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		// put '.' if it is not a start or end of a word, current letter is an uppercase letter,
		// and previous letter is a lowercase letter (cases like: "UserName"), or next letter is
		// also a lowercase letter and previous letter is not "_".
		if i > 0 && i < len(s)-1 && unicode.IsUpper(r) &&
			(unicode.IsLower(rune(s[i-1])) ||
				unicode.IsLower(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1]))) {
			b.WriteString("_")
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

// init initializes the parser parsing state. it scans the fields
// in a breath-first-search order and for each one of the field calls parseField.
func (p *Parser) init() error {
	t := indirect(reflect.TypeOf(p.Model))
	l := list.New()
	for i := 0; i < t.NumField(); i++ {
		l.PushFront(t.Field(i))
	}
	for l.Len() > 0 {
		f := l.Remove(l.Front()).(reflect.StructField)
		_, ok := f.Tag.Lookup(p.TagName)
		switch t := indirect(f.Type); {
		// no matter what the type of this field. if it has a tag,
		// it is probably a filterable or sortable.
		case ok:
			if err := p.parseField(f); err != nil {
				return err
			}
		case t.Kind() == reflect.Struct:
			for i := 0; i < t.NumField(); i++ {
				f1 := t.Field(i)
				if !f.Anonymous {
					f1.Name = f.Name + p.FieldSep + f1.Name
				}
				l.PushFront(f1)
			}
		case f.Anonymous:
			p.Log("ignore embedded field %q that is not struct type", f.Name)
		}
	}
	return nil
}

// parseField parses the given struct field tag, and add a rule
// in the parser according to its type and the options that were set on the tag.
func (p *Parser) parseField(sf reflect.StructField) error {
	f := &field{
		Column:    p.ColumnFn(sf.Name),
		Alias:     p.ColumnFn(sf.Name),
		CovertFn:  valueFn,
		FilterOps: make(map[string]bool),
	}
	layout := time.RFC3339
	opts := strings.Split(sf.Tag.Get(p.TagName), ",")
	datatype := ""
	for _, opt := range opts {
		switch s := strings.TrimSpace(opt); {
		case s == "sort":
			f.Sortable = true
		case s == "filter":
			f.Filterable = true
		case strings.HasPrefix(opt, "column"):
			f.Column = strings.TrimPrefix(opt, "column=")
		case strings.HasPrefix(opt, "alias"):
			f.Alias = strings.TrimPrefix(opt, "alias=")
		case strings.HasPrefix(opt, "datatype"):
			datatype = strings.TrimPrefix(opt, "datatype=")
		case strings.HasPrefix(opt, "layout"):
			layout = strings.TrimPrefix(opt, "layout=")
			// if it's one of the standard layouts, like: RFC822 or Kitchen.
			if ly, ok := layouts[layout]; ok {
				layout = ly
			}
			// test the layout on a value (on itself). however, some layouts are invalid
			// time values for time.Parse, due to formats such as _ for space padding and
			// Z for zone information.
			v := strings.NewReplacer("_", " ", "Z", "+").Replace(layout)
			if _, err := time.Parse(layout, v); err != nil {
				return fmt.Errorf("rql: layout %q is not parsable: %v", layout, err)
			}
		default:
			p.Log("Ignoring unknown option %q in struct tag", opt)
		}
	}
	var filterOps []Op
	switch typ := indirect(sf.Type); typ.Kind() {
	case reflect.Bool:
		f.ValidateFn = validateBool
		filterOps = append(filterOps, EQ, NEQ)
	case reflect.String:
		f.ValidateFn = validateString
		filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE, LIKE, NLIKE, ILIKE, NILIKE)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f.ValidateFn = validateInt
		f.CovertFn = convertInt
		filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		f.ValidateFn = validateUInt
		f.CovertFn = convertInt
		filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
	case reflect.Float32, reflect.Float64:
		f.ValidateFn = validateFloat
		filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
	case reflect.Struct:
		switch v := reflect.Zero(typ); v.Interface().(type) {
		case sql.NullBool:
			f.ValidateFn = validateBool
			filterOps = append(filterOps, EQ, NEQ)
		case sql.NullString:
			f.ValidateFn = validateString
			filterOps = append(filterOps, EQ, NEQ)
		case sql.NullInt64:
			f.ValidateFn = validateInt
			f.CovertFn = convertInt
			filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
		case sql.NullFloat64:
			f.ValidateFn = validateFloat
			filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
		case time.Time:
			f.ValidateFn = validateTime(layout)
			f.CovertFn = convertTime(layout)
			filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
		default:
			if v.Type().ConvertibleTo(reflect.TypeOf(time.Time{})) {
				f.ValidateFn = validateTime(layout)
				f.CovertFn = convertTime(layout)
				filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
			} else {
				switch datatype {
				case "bool":
					f.ValidateFn = validateBool
					filterOps = append(filterOps, EQ, NEQ)
				case "string":
					f.ValidateFn = validateString
					filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE, LIKE, NLIKE, ILIKE, NILIKE)
				case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
					f.ValidateFn = validateInt
					f.CovertFn = convertInt
					filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
				case "float32", "float64":
					f.ValidateFn = validateFloat
					filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
				case "time.Time":
					f.ValidateFn = validateTime(layout)
					f.CovertFn = convertTime(layout)
					filterOps = append(filterOps, EQ, NEQ, LT, LTE, GT, GTE)
				default:
					return fmt.Errorf("rql: field type for %q is not supported", sf.Name)
				}
			}
		}
	default:
		return fmt.Errorf("rql: field type for %q is not supported", sf.Name)
	}
	for _, op := range filterOps {
		f.FilterOps[p.op(op)] = true
	}
	p.fields[f.Alias] = f
	return nil
}

type parseState struct {
	*Parser               // reference of the parser config
	values  []interface{} // query values
	Exp     []exp.Expression
}

var parseStatePool sync.Pool

func (p *Parser) newParseState() (ps *parseState) {
	if v := parseStatePool.Get(); v != nil {
		ps = v.(*parseState)
		ps.Exp = make([]exp.Expression, 0, 5)
		ps.values = nil
	} else {
		ps = new(parseState)
		// currently we're using an arbitrary size as the capacity of initial buffer.
		// What we can do in the future is to track the size of parse results, and use
		// the average value. Same thing applies to the `values` field below.
		ps.Exp = make([]exp.Expression, 0, 5)
	}
	ps.values = make([]interface{}, 0, 8)
	ps.Parser = p
	return
}

// sort build the sort clause.
func (p *Parser) sort(fields []string) []exp.OrderedExpression {
	sortParams := make([]exp.OrderedExpression, len(fields))
	for i, field := range fields {
		expect(field != "", "sort field can not be empty")
		orderBy := Asc
		// if the sort field prefixed by an order indicator.
		if order, ok := sortDirection[field[0]]; ok {
			orderBy = order
			field = field[1:]
		}
		expect(p.fields[field] != nil, "unrecognized key %q for sorting", field)
		expect(p.fields[field].Sortable, "field %q is not sortable", field)
		colName := p.colName(p.fields[field].Column)
		col := goqu.I(colName)
		if orderBy == Asc {
			sortParams[i] = col.Asc()
		} else {
			sortParams[i] = col.Desc()
		}
	}
	return sortParams
}

func (p *parseState) and(f map[string]interface{}) exp.ExpressionList {
	andExp := []exp.Expression{}
	for k, v := range f {
		switch {
		case k == p.op(OR):
			terms, ok := v.([]interface{})
			expect(ok, "$or must be type array")
			andExp = append(andExp, p.relOp(OR, terms))
		case k == p.op(AND):
			terms, ok := v.([]interface{})
			expect(ok, "$and must be type array")
			andExp = append(andExp, p.relOp(AND, terms))
		case p.fields[k] != nil:
			expect(p.fields[k].Filterable, "field %q is not filterable", k)
			andExp = append(andExp, p.field(p.fields[k], v))
		default:
			expect(false, "unrecognized key %q for filtering", k)
		}
	}
	return goqu.And(andExp...)
}

func (p *parseState) relOp(op Op, terms []interface{}) exp.ExpressionList {
	exp := []exp.Expression{}
	for _, t := range terms {
		mt, ok := t.(map[string]interface{})
		expect(ok, "expressions for $%s operator must be type object", op)
		exp = append(exp, p.and(mt))
	}

	if op == OR {
		return goqu.Or(exp...)
	}
	return goqu.And(exp...)
}

func (p *parseState) field(f *field, v interface{}) exp.ExpressionList {
	and := []exp.Expression{}

	terms, ok := v.(map[string]interface{})
	// default equality check.
	if !ok {
		//must(f.ValidateFn(v), "invalid datatype for field %q", f.Name)
		value := f.CovertFn(v)
		and = append(and, p.fmtOp(f.Column, EQ, value))
		p.values = append(p.values, value)
	}

	for opName, opVal := range terms {
		expect(f.FilterOps[opName], "can not apply op %q on field %q", opName, f.Alias)
		//must(f.ValidateFn(opVal), "invalid datatype or format for field %q", f.Name)

		value := f.CovertFn(opVal)
		and = append(and, p.fmtOp(f.Column, Op(opName[1:]), value))
		p.values = append(p.values, value)
	}
	return goqu.And(and...)
}

func (p *Parser) fmtOp(field string, op Op, value interface{}) exp.Expression {
	col := goqu.C(p.colName(field))

	switch op {
	case EQ:
		return col.Eq(value)
	case NEQ:
		return col.Neq(value)
	case GT:
		return col.Gt(value)
	case GTE:
		return col.Gte(value)
	case LT:
		return col.Lt(value)
	case LTE:
		return col.Lte(value)
	case LIKE:
		return col.Like(value)
	case NLIKE:
		return col.NotLike(value)
	case ILIKE:
		return col.ILike(value)
	case NILIKE:
		return col.NotILike(value)
	}
	return nil
}

// colName formats the query field to database column name in cases the user configured a custom
// field separator. for example: if the user configured the field separator to be ".", the fields
// like "address.name" will be changed to "address_name".
func (p *Parser) colName(field string) string {
	if p.FieldSep != DefaultFieldSep {
		return strings.Replace(field, p.FieldSep, DefaultFieldSep, -1)
	}
	return field
}

func (p *Parser) op(op Op) string {
	return p.OpPrefix + string(op)
}

// expect panic if the condition is false.
func expect(cond bool, msg string, args ...interface{}) {
	if !cond {
		panic(&ParseError{fmt.Sprintf(msg, args...)})
	}
}

// must panics if the error is not nil.
func must(err error, msg string, args ...interface{}) {
	if err != nil {
		args = append(args, err)
		panic(&ParseError{fmt.Sprintf(msg+": %s", args...)})
	}
}

// indirect returns the item at the end of indirection.
func indirect(t reflect.Type) reflect.Type {
	for ; t.Kind() == reflect.Ptr; t = t.Elem() {
	}
	return t
}

// --------------------------------------------------------
// Validators and Converters

func errorType(v interface{}, expected string) error {
	actual := "nil"
	if v != nil {
		actual = reflect.TypeOf(v).Kind().String()
	}
	return fmt.Errorf("expect <%s>, got <%s>", expected, actual)
}

// validate that the underlined element of given interface is a boolean.
func validateBool(v interface{}) error {
	if _, ok := v.(bool); !ok {
		return errorType(v, "bool")
	}
	return nil
}

// validate that the underlined element of given interface is a string.
func validateString(v interface{}) error {
	if _, ok := v.(string); !ok {
		return errorType(v, "string")
	}
	return nil
}

// validate that the underlined element of given interface is a float.
func validateFloat(v interface{}) error {
	if _, ok := v.(float64); !ok {
		return errorType(v, "float64")
	}
	return nil
}

// validate that the underlined element of given interface is an int.
func validateInt(v interface{}) error {
	n, ok := v.(float64)
	if !ok {
		return errorType(v, "int")
	}
	if math.Trunc(n) != n {
		return errors.New("not an integer")
	}
	return nil
}

// validate that the underlined element of given interface is an int and greater than 0.
func validateUInt(v interface{}) error {
	if err := validateInt(v); err != nil {
		return err
	}
	if v.(float64) < 0 {
		return errors.New("not an unsigned integer")
	}
	return nil
}

// validate that the underlined element of this interface is a "datetime" string.
func validateTime(layout string) func(interface{}) error {
	return func(v interface{}) error {
		s, ok := v.(string)
		if !ok {
			return errorType(v, "string")
		}
		_, err := time.Parse(layout, s)
		return err
	}
}

// convert float to int.
func convertInt(v interface{}) interface{} {
	return int(v.(float64))
}

// convert string to time object.
func convertTime(layout string) func(interface{}) interface{} {
	return func(v interface{}) interface{} {
		t, _ := time.Parse(layout, v.(string))
		return t
	}
}

// nop converter.
func valueFn(v interface{}) interface{} {
	return v
}

// layouts holds all standard time.Time layouts.
var layouts = map[string]string{
	"ANSIC":       time.ANSIC,
	"UnixDate":    time.UnixDate,
	"RubyDate":    time.RubyDate,
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Kitchen":     time.Kitchen,
	"Stamp":       time.Stamp,
	"StampMilli":  time.StampMilli,
	"StampMicro":  time.StampMicro,
	"StampNano":   time.StampNano,
}
