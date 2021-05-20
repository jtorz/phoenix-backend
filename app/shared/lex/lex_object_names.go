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
	FndModule        string
	FndAction        string
	FndRole          string
	FndPrivilege     string
	FndUserRole      string
	FndUser          string
	FndPassword      string
	FndAccessAccount string
}{
	FndModule:        "fnd_module",
	FndAction:        "fnd_action",
	FndRole:          "fnd_role",
	FndPrivilege:     "fnd_privilege",
	FndUserRole:      "fnd_ser_role",
	FndUser:          "fnd_user",
	FndPassword:      "fnd_password",
	FndAccessAccount: "fnd_access_account",
}

// V database view names.
var V = struct {
	FndVPrivilegeRole string
}{
	FndVPrivilegeRole: "fnd_v_privilege_role",
}
