// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package {{.PackageName}}


var (	{{range .Tables }}
	// {{.GoCase}} column names for object {{.Name}}.
	{{.GoCase}} = struct { {{range .Columns }}
			{{.GoCase}} string {{end}}
	}{ {{range .Columns }}
			{{.GoCase}}: "{{.Name}}", {{end}}
	}
	{{end}}
)