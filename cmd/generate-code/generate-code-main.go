package main

import (
	"bytes"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"go/format"
	"log"
	"strings"
	"text/template"

	codegenasset "github.com/jtorz/phoenix-backend/app/assets/tools/code-generator-asset"
	"github.com/jtorz/phoenix-backend/utils/codegen"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ComponentType string

const (
	Model       ComponentType = "M"
	Dao         ComponentType = "D"
	Business    ComponentType = "B"
	HttpHandler ComponentType = "H"
	RestTest    ComponentType = "R"
	All         ComponentType = "A"
)

type TemplateData struct {
	ServiceName string // Service name in GoCase (Example: Foundation)
	ServiceAbbr string // Service abbreviation in GoCase (Example: Fnd)
	Entity      *codegen.Entity
}

func init() {
	flag.String("db", "", "REQUIRED - Database connection string")
	flag.String("schema", "", "REQUIRED - Database schema")
	flag.String("table", "", "REQUIRED - Database table (Example: fnduser)")
	flag.String("svc", "", "REQUIRED - Service name (Example: foundation)")
	flag.String("svcAbbr", "", "REQUIRED - Service abbreviation (Example: fnd)")
	flag.String("component", "A", "REQUIRED - Type of component to generate ('M:Model' 'D:Dao' 'B:Business' 'H:HttpHandler' 'R:RestTest' 'A:All')")
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

func getComponent(key string) ComponentType {
	v := viper.GetString(key)
	if v == "" {
		v = "A"
	}
	switch c := ComponentType(v); c {
	case Model, Dao, Business, HttpHandler, RestTest, All:
		return c
	default:
		log.Fatalf("uknown component generation (%s)", v)
		return All
	}
}

type Components map[ComponentType]Component

type Component struct {
	Template string
	Out      string
}

func main() {
	db := loadDB()
	schema := getReqFlag("schema")
	table := getReqFlag("table")
	entity, err := codegen.NewEntity(context.Background(), db, schema, table)
	if err != nil {
		log.Fatal(err)
	}
	templateData := TemplateData{
		ServiceName: getReqFlag("svc"),
		ServiceAbbr: getReqFlag("svcAbbr"),
		Entity:      entity,
	}
	Components := Components{
		Model: Component{
			Out:      "app/services/%s/%smodel/%s_m_%s.go",
			Template: codegenasset.ModelTPL,
		},
		/*Dao: Component{
			Out:      "app/services/%s/%smodel/%s_m_%s.go",
			Template: codegenasset.ModelTPL,
		},
		Business: Component{
			Out:      "app/services/%s/%smodel/%s_m_%s.go",
			Template: codegenasset.ModelTPL,
		},
		HttpHandler: Component{
			Out:      "app/services/%s/%smodel/%s_m_%s.go",
			Template: codegenasset.ModelTPL,
		},
		RestTest: Component{
			Out:      "app/services/%s/%smodel/%s_m_%s.go",
			Template: codegenasset.ModelTPL,
		},*/
	}

	for _, tpl := range Components {
		var file []byte
		file, err = executeTemplate(tpl.Template, templateData)
		if err != nil {
			log.Fatal("objects", err)
		}
		fmt.Println(string(file))
		/*err = ioutil.WriteFile(gen.OutputDir+tpl.ouput, file, 0644)
		if err != nil {
			log.Fatal(err)
		}*/
	}
}

func formatOut(out, serviceAbbr, entityName string) string {
	return fmt.Sprintf(out, serviceAbbr, serviceAbbr, serviceAbbr, entityName)
}

func loadDB() *sql.DB {
	var err error
	dbcon := getReqFlag("db")
	db, err := sql.Open("postgres", dbcon)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func executeTemplate(textTemplate string, templateData TemplateData) ([]byte, error) {
	tpl, err := template.New("tablenames").
		Funcs(template.FuncMap{
			"lowercase": strings.ToLower,
		}).
		Parse(textTemplate)
	if err != nil {
		return nil, err
	}
	sb := bytes.Buffer{}
	tpl.Execute(&sb, templateData)
	//return sb.Bytes(), nil
	content, err := format.Source(sb.Bytes())
	if err != nil {
		return nil, err
	}
	return content, err
}
