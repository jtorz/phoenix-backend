// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.PackageName}}

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

// T database table names.
var T = struct { {{range .Tables}}
	{{.DBGoCase}} string {{end}}
}{ {{range .Tables}}
	{{.DBGoCase}} :"{{.DBName}}", {{end}}
}

// V database view names.
var V = struct { {{range .Views}}
	{{.DBGoCase}} string {{end}}
}{ {{range .Views}}
	{{.DBGoCase}} :"{{.DBName}}", {{end}}
}


{{range $T := .Tables }}
	{{range $FK := .Fks}}
	// {{$FK.DBGoCase}} returns the join expression for the foreign key from {{$T.DBGoCase}} to {{$FK.FTableGoCase}}.
	func {{$FK.DBGoCase}}(exps ...exp.Expression) exp.JoinCondition{
		exps = append(exps, goqu.Ex{ {{range $fkCol := $FK.Columns}}
				{{$T.DBGoCase}}.{{$fkCol.O.DBGoCase}}: goqu.I({{$FK.FTableGoCase}}.{{$fkCol.D.DBGoCase}}), {{end}}
		})
		return goqu.On(exps...)
	}
	{{end}}
{{end}}
