// Package {{.PackageName}} contains the dictionary (lexicon) of the database.
//
// The elements in the package are:
//
// * Table names
// * Table columns
// * View names
// * View columns
package {{.PackageName}}

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
