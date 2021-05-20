package lexasset

import _ "embed"

//go:embed object_names.tpl
var ObjectNamesTpl string

//go:embed object_column_names.tpl
var ObjectColumnNamesTpl string
