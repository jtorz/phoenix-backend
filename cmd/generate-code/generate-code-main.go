package main

import (
	"bytes"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	codegenasset "github.com/jtorz/phoenix-backend/app/assets/tools/code-generator-asset"
	"github.com/jtorz/phoenix-backend/utils/codegen"
	"github.com/jtorz/phoenix-backend/utils/stringset"
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

func (c ComponentType) String() string {
	switch c {
	case Model:
		return "Model"
	case Dao:
		return "Dao"
	case Business:
		return "Business"
	case HttpHandler:
		return "HttpHandler"
	case RestTest:
		return "RestTest"
	case All:
		return "All"
	default:
		return "????"
	}
}

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
	flag.Bool("write", true, "Should write to output files in the service directory.")
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

func getSelectedComponent() ComponentType {
	v := viper.GetString("component")
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

type Components []Component

type Component struct {
	Type     ComponentType
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
		Component{
			Type:     Model,
			Out:      "app/services/%s/%smodel/%s_m_%s.go",
			Template: codegenasset.ModelTPL,
		},
		Component{
			Type:     Dao,
			Out:      "app/services/%s/%sdao/%s_dao_%s.go",
			Template: codegenasset.DaoTPL,
		},
		Component{
			Type:     Business,
			Out:      "app/services/%s/%sbiz/%s_biz_%s.go",
			Template: codegenasset.BusinessTPL,
		},
		Component{
			Type:     HttpHandler,
			Out:      "app/services/%s/%shttp/%s_http_%s.go",
			Template: codegenasset.HandlerTPL,
		},
		Component{
			Type:     RestTest,
			Out:      "app/services/%s/%srest_test/%s_test_%s.rest",
			Template: codegenasset.RestTestTPL,
		},
	}

	isAppendFile := false
	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()
	comp := getSelectedComponent()
	for _, tpl := range Components {
		if comp != All && comp != tpl.Type {
			continue
		}
		var fileData []byte
		fileData, err = executeTemplate(tpl.Template, templateData, tpl.Type)
		if err != nil {
			return
		}

		outFileName := formatOut(tpl.Out, templateData.ServiceAbbr, *templateData.Entity)
		if !fileExists(outFileName) && viper.GetBool("write") {
			err = ioutil.WriteFile(outFileName, fileData, 0644)
			if err != nil {
				return
			}
		} else {
			fmt.Printf("Component not written file already exists or ignored due to flag 'write'. Check codegen.gotpl file. %s (%s)\n", tpl.Type, outFileName)
		}
		header := buildHeader(outFileName, tpl.Type)
		fileData = append(header, fileData...)
		if isAppendFile {
			err = appendFile("codegen.gotpl", fileData)
			if err != nil {
				return
			}
			continue
		} else {
			err = ioutil.WriteFile("codegen.gotpl", fileData, 0644)
			if err != nil {
				return
			}
		}

		isAppendFile = true
	}
}

func buildHeader(outFileName string, componentType ComponentType) []byte {
	name := " " + componentType.String() + " Component "
	s := (100 - len(name)) / 2
	buf := bytes.Buffer{}
	buf.Grow(550) // aproximate size

	buf.WriteRune('\n')
	buf.Write(bytes.Repeat([]byte("-"), 100))
	buf.WriteRune('\n')
	buf.Write(bytes.Repeat([]byte("-"), 100))
	buf.WriteRune('\n')

	buf.Write(bytes.Repeat([]byte("-"), s))
	buf.WriteString(name)
	buf.Write(bytes.Repeat([]byte("-"), s))
	if buf.Len() == 302 {
		buf.WriteRune('-')
	}
	buf.WriteRune('\n')

	buf.Write(bytes.Repeat([]byte("-"), 100))
	buf.WriteRune('\n')
	buf.Write(bytes.Repeat([]byte("-"), 100))
	buf.WriteRune('\n')
	buf.WriteRune('\n')

	buf.WriteString("outFile: " + outFileName)
	buf.WriteRune('\n')
	return buf.Bytes()
}

func appendFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(data); err != nil {
		return err
	}
	return nil
}

func formatOut(out, serviceAbbr string, entity codegen.Entity) string {
	serviceAbbr = strings.ToLower(serviceAbbr)
	return fmt.Sprintf(out, serviceAbbr, serviceAbbr, serviceAbbr, strings.ToLower(entity.DBName[4:]))
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

func executeTemplate(textTemplate string, templateData TemplateData, componentType ComponentType) ([]byte, error) {
	tpl, err := template.New("tablenames").
		Funcs(template.FuncMap{
			"lowercase":  strings.ToLower,
			"lowerfirst": stringset.LowerFirst,
			"upperfirst": stringset.UpperFirst,
		}).
		Parse(textTemplate)
	if err != nil {
		return nil, err
	}
	sb := bytes.Buffer{}
	tpl.Execute(&sb, templateData)
	if componentType == RestTest {
		return sb.Bytes(), nil
	}
	content, err := format.Source(sb.Bytes())
	if err != nil {
		return nil, err
	}
	return content, err
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
