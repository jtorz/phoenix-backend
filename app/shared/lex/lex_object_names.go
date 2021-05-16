// Package lex contains the dictionary (lexicon) of the database.
//
// The elements in the package are:
//
// * Table names
// * Table columns
// * view names
// * view columns
package lex

// T database table names.
var T = struct {
	Fndccomponent     string
	Fndtmodule        string
	Fndtaction        string
	Fndtrole          string
	Fndtprivilege     string
	FndtuserRole      string
	Fndtuser          string
	Fndtpassword      string
	FndcaccessType    string
	FndtaccessAccount string
	Emacsender        string
	EmactemplateType  string
	Ematheader        string
	Ematfooter        string
	Emattemplate      string
	Emabrecord        string
}{
	Fndccomponent:     "fndccomponent",
	Fndtmodule:        "fndtmodule",
	Fndtaction:        "fndtaction",
	Fndtrole:          "fndtrole",
	Fndtprivilege:     "fndtprivilege",
	FndtuserRole:      "fnduser_role",
	Fndtuser:          "fndtuser",
	Fndtpassword:      "fndtpassword",
	FndcaccessType:    "fndcaccess_type",
	FndtaccessAccount: "fndtaccess_account",
	Emacsender:        "emacsender",
	EmactemplateType:  "emactemplate_type",
	Ematheader:        "ematheader",
	Ematfooter:        "ematfooter",
	Emattemplate:      "emattemplate",
	Emabrecord:        "emabrecord",
}

// V database view names.
var V = struct {
	FndvprivilegeRole string
}{
	FndvprivilegeRole: "fndvprivilege_role",
}
