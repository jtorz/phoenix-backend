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
}

type TplColumn struct {
	GoCase   string
	Name     string
	Nullable string
	DataType string
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

func main() {
	var objects, objectNames, testFile []byte
	gen := Generator{}
	gen.TPLData.PackageName = getReqFlag("pkg")
	gen.TPLData.TestPackage = getReqFlag("testPkg")
	gen.OutputDir = getReqFlag("out")
	gen.Schema = getReqFlag("schema")
	gen.FilterPrefix = viper.GetString("filterPrefix")
	gen.checkAvoidOverwrite()
	gen.loadDB()
	err := gen.getObjects()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(gen)
	//os.Exit(0)
	objects, err = gen.executeTemplate(lexasset.ObjectNamesTpl)
	if err != nil {
		log.Fatal("objects", err)
	}
	objectNames, err = gen.executeTemplate(lexasset.ObjectColumnNamesTpl)
	if err != nil {
		log.Fatal("objectNames", err)
	}

	testFile, err = gen.executeTemplate(lexasset.TestTpl)
	if err != nil {
		log.Fatal("objectNames", err)
	}
	file := gen.OutputDir + "/lex_object_names.go"
	err = ioutil.WriteFile(file, objects, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file = gen.OutputDir + "/lex_object_columns.go"
	err = ioutil.WriteFile(file, objectNames, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file = gen.OutputDir + "/lex_test.go"
	err = ioutil.WriteFile(file, testFile, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (gen *Generator) checkAvoidOverwrite() {
	if viper.GetBool("overwrite") {
		return
	}
	file := gen.OutputDir + "/lex_object_names.go"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf("File %s already exists. Use --overwrite flag or delete file to continue.", file)
	}
	file = gen.OutputDir + "/lex_object_columns.go"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf("File %s already exists. Use --overwrite flag or delete file to continue.", file)
	}
	file = gen.OutputDir + "/lex_test.go"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf("File %s already exists. Use --overwrite flag or delete file to continue.", file)
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
