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
	Fndccomponent:     "fnd_component",
	Fndtmodule:        "fnd_module",
	Fndtaction:        "fnd_action",
	Fndtrole:          "fnd_role",
	Fndtprivilege:     "fnd_privilege",
	FndtuserRole:      "fnd_ser_role",
	Fndtuser:          "fnd_user",
	Fndtpassword:      "fnd_password",
	FndcaccessType:    "fnd_access_type",
	FndtaccessAccount: "fnd_access_account",
	Emacsender:        "ema_sender",
	EmactemplateType:  "ema_template_type",
	Ematheader:        "ema_header",
	Ematfooter:        "ema_footer",
	Emattemplate:      "ema_template",
	Emabrecord:        "ema_record",
}

// V database view names.
var V = struct {
	FndVPrivilegeRole string
}{
	FndVPrivilegeRole: "fnd_v_privilege_role",
}
