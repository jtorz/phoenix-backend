// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.PackageName}}

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

// T database table names.
var T = struct { {{range .Tables}}
	{{.GoCase}} string {{end}}
}{ {{range .Tables}}
	{{.GoCase}} :"{{.Name}}", {{end}}
}

// V database view names.
var V = struct { {{range .Views}}
	{{.GoCase}} string {{end}}
}{ {{range .Views}}
	{{.GoCase}} :"{{.Name}}", {{end}}
}



{{range $T := .Tables }}
	{{range $FK := .Fks}}
	// {{$FK.ConstraintName}} returns the join expression for the foreign key from {{$T.GoCase}} to {{$FK.FTable}}.
	func {{$FK.ConstraintName}}(exps ...exp.Expression) exp.JoinCondition{
		exps = append(exps, goqu.Ex{ {{range $fkCol := $FK.Columns}}
				{{$T.GoCase}}.{{$fkCol.Orig}}: goqu.I({{$FK.FTable}}.{{$fkCol.Dest}}), {{end}}
		})
		return goqu.On(exps...)
	}
	{{end}}
{{end}}

