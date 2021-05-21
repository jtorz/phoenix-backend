// Package lex contains the dictionary (lexicon) of the database.
//
// The elements in the package are:
//
// lex_object_names.go
//   * Table names
//   * View names
//   * FK Constraints join expressions
//
//  lex_object_columns.go
//   * Table columns
//   * View columns
//
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package lex

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

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

// FndModuleFKFNDModule returns the join expression for the foreign key from FndModule to FndModule.
func FndModuleFKFNDModule(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndModule.ModParentID: goqu.I(FndModule.ModID),
	})
	return goqu.On(exps...)
}

// FndtactionFKFNDModule returns the join expression for the foreign key from FndAction to FndModule.
func FndtactionFKFNDModule(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndAction.ActModuleID: goqu.I(FndModule.ModID),
	})
	return goqu.On(exps...)
}

// FndtprivilegeFKFndtaction returns the join expression for the foreign key from FndPrivilege to FndAction.
func FndtprivilegeFKFndtaction(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndPrivilege.PriActionID: goqu.I(FndAction.ActActionID),
		FndPrivilege.PriModuleID: goqu.I(FndAction.ActModuleID),
	})
	return goqu.On(exps...)
}

// FndtprivilegeFKFndtrole returns the join expression for the foreign key from FndPrivilege to FndRole.
func FndtprivilegeFKFndtrole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndPrivilege.PriRoleID: goqu.I(FndRole.RolID),
	})
	return goqu.On(exps...)
}

// FndtpasswordFKFndtuser returns the join expression for the foreign key from FndPassword to FndUser.
func FndtpasswordFKFndtuser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndPassword.PasUserID: goqu.I(FndUser.UseID),
	})
	return goqu.On(exps...)
}

// FndtuserRoleFKFndtrole returns the join expression for the foreign key from FndUserRole to FndRole.
func FndtuserRoleFKFndtrole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndUserRole.UsrRoleID: goqu.I(FndRole.RolID),
	})
	return goqu.On(exps...)
}

// FndtuserRoleFKFndtuser returns the join expression for the foreign key from FndUserRole to FndUser.
func FndtuserRoleFKFndtuser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndUserRole.UsrUserID: goqu.I(FndUser.UseID),
	})
	return goqu.On(exps...)
}

// FndtroleNavigatorFKFndtnavigator returns the join expression for the foreign key from FndRoleNavigator to FndNavigator.
func FndtroleNavigatorFKFndtnavigator(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndRoleNavigator.RonNavigatorID: goqu.I(FndNavigator.NavID),
	})
	return goqu.On(exps...)
}

// FndtroleNavigatorFKFndtrole returns the join expression for the foreign key from FndRoleNavigator to FndRole.
func FndtroleNavigatorFKFndtrole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndRoleNavigator.RonRoleID: goqu.I(FndRole.RolID),
	})
	return goqu.On(exps...)
}

// FndtaccessAccountFKFndtuser returns the join expression for the foreign key from FndAccountAccess to FndUser.
func FndtaccessAccountFKFndtuser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndAccountAccess.AcaUserID: goqu.I(FndUser.UseID),
	})
	return goqu.On(exps...)
}
