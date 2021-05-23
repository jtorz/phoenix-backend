// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.PackageName}}


{{range .Tables }}
// Table{{.GoCase}} column names for table {{.Name}}.
type Table{{.GoCase}} struct { {{range .Columns }}
	{{.GoCase}} string `database:"{{.Nullable}},{{.DataType}}"` {{end}}
}

var {{.GoCase}} = Table{{.GoCase}}{ {{range .Columns }}
	{{.GoCase}}: "{{.Name}}", {{end}}
}
{{end}}

{{range .Views }}
// View{{.GoCase}} column names for view {{.Name}}.
type View{{.GoCase}} struct { {{range .Columns }}
	{{.GoCase}} string `database:"{{.Nullable}},{{.DataType}}"` {{end}}
}

var {{.GoCase}} = View{{.GoCase}}{ {{range .Columns }}
	{{.GoCase}}: "{{.Name}}", {{end}}
}
{{end}}
