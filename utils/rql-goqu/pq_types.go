package rqlgq

// TranslateColumnType converts postgres database types to Go types, for example
// "varchar" to "string" and "bigint" to "int64". It returns this parsed data
// as a Column object.
func TranslateColumnType(DBType string) string {
	switch DBType {
	case "bigint", "bigserial":
		return "int64"
	case "integer", "serial":
		return "int"
	case "oid":
		return "uint32"
	case "smallint", "smallserial":
		return "int16"
	case "decimal", "numeric":
		return "float64"
	case "double precision":
		return "float64"
	case "real":
		return "float32"
	case "bit", "interval", "uuint", "bit varying", "character", "money", "character varying", "cidr", "inet", "macaddr", "text", "uuid", "xml":
		return "string"
	case `"char"`:
		return "string"
	case "boolean":
		return "bool"
	case "date", "time", "timestamp without time zone", "timestamp with time zone", "time without time zone", "time with time zone":
		return "time.Time"
	default:
		return "string"
	}
}
