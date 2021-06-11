// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dalhelper

import (
	"database/sql"
	"os"
	"reflect"
	"strings"
	"testing"

	configtest "github.com/jtorz/phoenix-backend/app/config/testconfig"
)

var mainDB *sql.DB

func TestMain(m *testing.M) {
	mainDB = configtest.MainDB()
	if mainDB != nil {
		os.Exit(m.Run())
	}
}

func TestTables(t *testing.T) {
	testSelect(t, T)
}

func TestViews(t *testing.T) {
	testSelect(t, V)
}

// testSelect executes a basic select statement for each table in the struct to check if the tables exists.
func testSelect(t *testing.T, objects interface{}) {
	val := reflect.ValueOf(objects)
	for i := 0; i < val.NumField(); i++ {
		goName := val.Type().Field(i).Name
		tableName := val.Field(i).Interface().(string)
		if err := selectAllFrom(mainDB, tableName); err != nil {
			t.Errorf(`MainDB dalhelper.T.%s: SELECT 1 FROM "%s": %s`, goName, tableName, err)
		}
	}
}

// selectAllFrom executes a basic select statement to check if the table exists.
func selectAllFrom(db *sql.DB, table string) error {
	_, err := db.Exec(`SELECT 1 FROM "` + table + `" LIMIT 1`)
	return err
}

func TestColumns(t *testing.T) {
	var testCases = []struct {
		tableName   string
		tableStruct interface{}
	}{
		{T.CoreAccountAccess, CoreAccountAccess},
		{T.CoreAction, CoreAction},
		{T.CoreActionRoute, CoreActionRoute},
		{T.CoreModule, CoreModule},
		{T.CoreNavElement, CoreNavElement},
		{T.CoreNavElementRole, CoreNavElementRole},
		{T.CorePassword, CorePassword},
		{T.CorePrivilege, CorePrivilege},
		{T.CoreRole, CoreRole},
		{T.CoreUser, CoreUser},
		{T.CoreUserRole, CoreUserRole},
		{T.MailBRecord, MailBRecord},
		{T.MailFooter, MailFooter},
		{T.MailHeader, MailHeader},
		{T.MailSender, MailSender},
		{T.MailTemplate, MailTemplate},
		{T.MailTemplateType, MailTemplateType},

		{V.CoreVPrivilegeRole, CoreVPrivilegeRole},
	}
	m := getObjectMap()
	for _, test := range testCases {
		m[test.tableName] = true
		if err := testColumns(t, mainDB, test.tableName, test.tableStruct); err != nil {
			t.Errorf(`MainDB select colums from %s : %s`, test.tableName, err)
		}
	}
	for name, usedInTestCase := range m {
		if !usedInTestCase {
			t.Errorf(`%s not added in test case`, name)
		}
	}
}

// getObjectMap returns a map with the name of the views an tables.
func getObjectMap() map[string]bool {
	m := make(map[string]bool)
	setObjectsMap(m, T)
	setObjectsMap(m, V)
	return m
}

func setObjectsMap(m map[string]bool, object interface{}) {
	val := reflect.ValueOf(object)
	for i := 0; i < val.NumField(); i++ {
		tableName := val.Field(i).Interface().(string)
		m[tableName] = false
	}
}

func testColumns(t *testing.T, db *sql.DB, table string, tableStruct interface{}) error {
	val := reflect.ValueOf(tableStruct)
	cols := []string{}
	for i := 0; i < val.NumField(); i++ {
		goName := val.Type().Field(i).Name
		colName, ok := val.Field(i).Interface().(string)
		if !ok {
			t.Errorf("%s.%s must be string", val.String(), goName)
			continue
		}
		cols = append(cols, `"`+colName+`"`)
	}
	_, err := db.Exec(`SELECT ` + strings.Join(cols, ",") + ` FROM "` + table + `" LIMIT 1`)
	return err
}
