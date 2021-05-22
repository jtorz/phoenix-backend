// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.PackageName}}


var (	{{range .Tables }}
	// {{.GoCase}} column names for table {{.Name}}.
	{{.GoCase}} = struct { {{range .Columns }}
			{{.GoCase}} string `database:"{{.Nullable}},{{.DataType}}"` {{end}}
	}{ {{range .Columns }}
			{{.GoCase}}: "{{.Name}}", {{end}}
	}
	{{end}}

	{{range .Views }}
	// {{.GoCase}} column names for view {{.Name}}.
	{{.GoCase}} = struct { {{range .Columns }}
			{{.GoCase}} string `database:"{{.Nullable}},{{.DataType}}"` {{end}}
	}{ {{range .Columns }}
			{{.GoCase}}: "{{.Name}}", {{end}}
	}
	{{end}}
)