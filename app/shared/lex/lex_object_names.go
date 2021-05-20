// Package lex contains the dictionary (lexicon) of the database.
//
// The elements in the package are:
//
// * Table names
// * Table columns
// * View names
// * View columns
package lex

// T database table names.
var T = struct {
	FndModule        string
	FndAction        string
	FndRole          string
	FndPrivilege     string
	FndUser          string
	FndPassword      string
	FndUserRole      string
	FndRoleNavigator string
	FndNavigator     string
	FndAccountAccess string
}{
	FndModule:        "fnd_module",
	FndAction:        "fnd_action",
	FndRole:          "fnd_role",
	FndPrivilege:     "fnd_privilege",
	FndUser:          "fnd_user",
	FndPassword:      "fnd_password",
	FndUserRole:      "fnd_user_role",
	FndRoleNavigator: "fnd_role_navigator",
	FndNavigator:     "fnd_navigator",
	FndAccountAccess: "fnd_account_access",
}

// V database view names.
var V = struct {
	FndVPrivilegeRole string
}{
	FndVPrivilegeRole: "fnd_v_privilege_role",
}
