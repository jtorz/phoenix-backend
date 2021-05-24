// Package lexasset uses go:embed to load the templates for lex generation.
package lexasset

import _ "embed"

//go:embed object_names.tpl
var ObjectNamesTpl string

//go:embed object_column_names.tpl
var ObjectColumnNamesTpl string

//go:embed test.tpl
var TestTpl string
