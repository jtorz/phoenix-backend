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

// FndModuleFkFndModulePadre returns the join expression for the foreign key from FndModule to FndModule.
func FndModuleFkFndModulePadre(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndModule.ModParentID: goqu.I(FndModule.ModID),
	})
	return goqu.On(exps...)
}

// FndActionFkFndModule returns the join expression for the foreign key from FndAction to FndModule.
func FndActionFkFndModule(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndAction.ActModuleID: goqu.I(FndModule.ModID),
	})
	return goqu.On(exps...)
}

// FndPrivilegeFkFndAction returns the join expression for the foreign key from FndPrivilege to FndAction.
func FndPrivilegeFkFndAction(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndPrivilege.PriActionID: goqu.I(FndAction.ActActionID),
		FndPrivilege.PriModuleID: goqu.I(FndAction.ActModuleID),
	})
	return goqu.On(exps...)
}

// FndPrivilegeFkFndRole returns the join expression for the foreign key from FndPrivilege to FndRole.
func FndPrivilegeFkFndRole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndPrivilege.PriRoleID: goqu.I(FndRole.RolID),
	})
	return goqu.On(exps...)
}

// FndPasswordFkFndUser returns the join expression for the foreign key from FndPassword to FndUser.
func FndPasswordFkFndUser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndPassword.PasUserID: goqu.I(FndUser.UseID),
	})
	return goqu.On(exps...)
}

// FndUserRoleFkFndRole returns the join expression for the foreign key from FndUserRole to FndRole.
func FndUserRoleFkFndRole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndUserRole.UsrRoleID: goqu.I(FndRole.RolID),
	})
	return goqu.On(exps...)
}

// FndUserRoleFkFndUser returns the join expression for the foreign key from FndUserRole to FndUser.
func FndUserRoleFkFndUser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndUserRole.UsrUserID: goqu.I(FndUser.UseID),
	})
	return goqu.On(exps...)
}

// FndRoleNavigatorFkFndNavigator returns the join expression for the foreign key from FndRoleNavigator to FndNavigator.
func FndRoleNavigatorFkFndNavigator(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndRoleNavigator.RonNavigatorID: goqu.I(FndNavigator.NavID),
	})
	return goqu.On(exps...)
}

// FndRoleNavigatorFkFndRole returns the join expression for the foreign key from FndRoleNavigator to FndRole.
func FndRoleNavigatorFkFndRole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndRoleNavigator.RonRoleID: goqu.I(FndRole.RolID),
	})
	return goqu.On(exps...)
}

// FndAccountAccessFkFndUser returns the join expression for the foreign key from FndAccountAccess to FndUser.
func FndAccountAccessFkFndUser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndAccountAccess.AcaUserID: goqu.I(FndUser.UseID),
	})
	return goqu.On(exps...)
}
