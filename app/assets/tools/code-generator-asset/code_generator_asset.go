// Package codegenasset uses go:embed to load the templates for code generation.
package codegenasset

import _ "embed"

//go:embed 01_model.tpl
var ModelTPL string

//go:embed 02_dao.tpl
var DaoTPL string

//go:embed 03_business.tpl
var BusinessTPL string

//go:embed 04_handler.tpl
var HandlerTPL string

//go:embed 05_rest_test.tpl
var RestTestTPL string
