// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.PackageName}}
{{define "NULLABLE"}}{{if . }}N{{else}}-{{end}}{{end}}

{{range .Tables }}
// Table{{.DBGoCase}} column names for table {{.DBName}}.
type Table{{.DBGoCase}} struct { {{range .Columns }}
	{{.DBGoCase}} string `database:"{{template "NULLABLE" .DBNullable}},datatype={{.DBDataType}}"`{{end}}
}

var {{.DBGoCase}} = Table{{.DBGoCase}}{ {{range .Columns }}
	{{.DBGoCase}}: "{{.DBName}}", {{end}}
}
{{end}}

{{range .Views }}
// View{{.DBGoCase}} column names for view {{.DBName}}.
type View{{.DBGoCase}} struct { {{range .Columns }}
	{{.DBGoCase}} string `database:"{{template "NULLABLE" .DBNullable}},datatype={{.DBDataType}}"` {{end}}
}

var {{.DBGoCase}} = View{{.DBGoCase}}{ {{range .Columns }}
	{{.DBGoCase}}: "{{.DBName}}", {{end}}
}
{{end}}
