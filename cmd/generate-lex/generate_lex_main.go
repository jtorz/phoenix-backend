package main

import (
	"bytes"
	"database/sql"
	"flag"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	lexasset "github.com/jtorz/phoenix-backend/app/assets/tools/lex-generator-asset"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	// postgres driver
	_ "github.com/lib/pq"
)

type Generator struct {
	DB           *sql.DB
	Schema       string
	OutputDir    string
	FilterPrefix string
	TPLData      TPLData
}

type TPLData struct {
	PackageName string
	TestPackage string
	Tables      []TplObject
	Views       []TplObject
}

type TplObject struct {
	GoCase  string
	Name    string
	Columns []TplColumn
	Fks     []Fk
}

type TplColumn struct {
	GoCase   string // database column name in Go case (Examples: UseID, UseName)
	Name     string // database column name (Examples: use_id, use_name)
	Nullable string // database value is nullable
	DataType string // database data type
	Field    string // Go field name (Examples: ID, Name)
}

type Fk struct {
	ConstraintName string
	Columns        []FKCol
	FTable         string
}
type FKCol struct {
	Orig, Dest string
}

func init() {
	flag.String("pkg", "lex", "go package name")
	flag.String("testPkg", "github.com/jtorz/phoenix-backend/app/config/configtest", "go test package name")
	flag.String("out", "app/shared/lex", "database connection string")
	flag.String("schema", "", "database schema")
	flag.String("db", "", "database connection string")
	flag.String("filterPrefix", "", "database connection string")
	flag.Bool("overwrite", false, "database connection string")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}
func getReqFlag(key string) string {
	v := viper.GetString(key)
	if v == "" {
		log.Fatal(key, " is required")
	}
	return v
}

type templateInfo struct {
	template *string
	ouput    string
}

func main() {
	templates := []templateInfo{
		{
			template: &lexasset.ObjectNamesTpl,
			ouput:    "/lex_object_names.go",
		},
		{
			template: &lexasset.ObjectColumnNamesTpl,
			ouput:    "/lex_object_columns.go",
		},
		{
			template: &lexasset.TestTpl,
			ouput:    "/lex_test.go",
		},
	}

	gen := Generator{}
	gen.TPLData.PackageName = getReqFlag("pkg")
	gen.TPLData.TestPackage = getReqFlag("testPkg")
	gen.OutputDir = getReqFlag("out")
	gen.Schema = getReqFlag("schema")
	gen.FilterPrefix = viper.GetString("filterPrefix")
	gen.checkAvoidOverwrite(templates)
	gen.loadDB()
	err := gen.getObjects()
	if err != nil {
		log.Fatal(err)
	}

	for _, tpl := range templates {
		var file []byte
		file, err = gen.executeTemplate(*tpl.template)
		if err != nil {
			log.Fatal("objects", err)
		}
		err = ioutil.WriteFile(gen.OutputDir+tpl.ouput, file, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (gen *Generator) checkAvoidOverwrite(templates []templateInfo) {
	if viper.GetBool("overwrite") {
		return
	}
	for _, tpl := range templates {
		file := gen.OutputDir + tpl.ouput
		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Fatalf("File %s already exists. Use --overwrite flag or delete file to continue.", file)
		}
	}
}

func (gen *Generator) loadDB() {
	var err error
	dbcon := getReqFlag("db")
	gen.DB, err = sql.Open("postgres", dbcon)
	if err != nil {
		log.Fatal(err)
	}
}

func (gen *Generator) executeTemplate(textTemplate string) ([]byte, error) {
	tpl, err := template.New("tablenames").Parse(textTemplate)
	if err != nil {
		return nil, err
	}
	sb := bytes.Buffer{}
	tpl.Execute(&sb, gen.TPLData)
	content, err := format.Source(sb.Bytes())
	if err != nil {
		return nil, err
	}
	return content, err
}
