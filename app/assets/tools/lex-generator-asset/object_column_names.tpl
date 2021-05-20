package {{.PackageName}}

var (
	{{range .Tables }}
	// {{.GoCase}} column names for object {{.Name}}.
	{{.GoCase}} = struct { {{range .Columns }}
			{{.GoCase}} string {{end}}
	}{ {{range .Columns }}
			{{.GoCase}}: {{.Name}}, {{end}}
	}
	{{end}}
)