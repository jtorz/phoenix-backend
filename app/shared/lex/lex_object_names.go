// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lex

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

// T database table names.
var T = struct {
	FndAccountAccess  string
	FndAction         string
	FndActionRoute    string
	FndModule         string
	FndNavElement     string
	FndNavElementRole string
	FndPassword       string
	FndPrivilege      string
	FndRole           string
	FndUser           string
	FndUserRole       string
	MailBRecord       string
	MailFooter        string
	MailHeader        string
	MailSender        string
	MailTemplate      string
	MailTemplateType  string
}{
	FndAccountAccess:  "fnd_account_access",
	FndAction:         "fnd_action",
	FndActionRoute:    "fnd_action_route",
	FndModule:         "fnd_module",
	FndNavElement:     "fnd_nav_element",
	FndNavElementRole: "fnd_nav_element_role",
	FndPassword:       "fnd_password",
	FndPrivilege:      "fnd_privilege",
	FndRole:           "fnd_role",
	FndUser:           "fnd_user",
	FndUserRole:       "fnd_user_role",
	MailBRecord:       "mail_b_record",
	MailFooter:        "mail_footer",
	MailHeader:        "mail_header",
	MailSender:        "mail_sender",
	MailTemplate:      "mail_template",
	MailTemplateType:  "mail_template_type",
}

// V database view names.
var V = struct {
	FndVPrivilegeRole string
}{
	FndVPrivilegeRole: "fnd_v_privilege_role",
}

// FndAccountAccessFkFndUser returns the join expression for the foreign key from FndAccountAccess to FndUser.
func FndAccountAccessFkFndUser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndAccountAccess.AcaUserID: goqu.I(FndUser.UseID),
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

// FndActionRouteFkFndAction returns the join expression for the foreign key from FndActionRoute to FndAction.
func FndActionRouteFkFndAction(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndActionRoute.AcrActionID: goqu.I(FndAction.ActActionID),
		FndActionRoute.AcrModuleID: goqu.I(FndAction.ActModuleID),
	})
	return goqu.On(exps...)
}

// FndModuleFkFndModulePadre returns the join expression for the foreign key from FndModule to FndModule.
func FndModuleFkFndModulePadre(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndModule.ModParentID: goqu.I(FndModule.ModID),
	})
	return goqu.On(exps...)
}

// FndNavElementFkFndNavElementParent returns the join expression for the foreign key from FndNavElement to FndNavElement.
func FndNavElementFkFndNavElementParent(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndNavElement.NaeParentID: goqu.I(FndNavElement.NaeID),
	})
	return goqu.On(exps...)
}

// FndNavElementRoleFkFndNavElement returns the join expression for the foreign key from FndNavElementRole to FndNavElement.
func FndNavElementRoleFkFndNavElement(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndNavElementRole.NerNavElementID: goqu.I(FndNavElement.NaeID),
	})
	return goqu.On(exps...)
}

// FndNavElementRoleFkFndRole returns the join expression for the foreign key from FndNavElementRole to FndRole.
func FndNavElementRoleFkFndRole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		FndNavElementRole.NerRoleID: goqu.I(FndRole.RolID),
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

// MailBRecordFkFndUser returns the join expression for the foreign key from MailBRecord to FndUser.
func MailBRecordFkFndUser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		MailBRecord.RecSenderUserID: goqu.I(FndUser.UseID),
	})
	return goqu.On(exps...)
}

// MailBRecordFkMailTemplateType returns the join expression for the foreign key from MailBRecord to MailTemplateType.
func MailBRecordFkMailTemplateType(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		MailBRecord.RecTypeID: goqu.I(MailTemplateType.TetID),
	})
	return goqu.On(exps...)
}

// MailTemplateFkMailFooter returns the join expression for the foreign key from MailTemplate to MailFooter.
func MailTemplateFkMailFooter(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		MailTemplate.TemFooterID: goqu.I(MailFooter.FooID),
	})
	return goqu.On(exps...)
}

// MailTemplateFkMailHeader returns the join expression for the foreign key from MailTemplate to MailHeader.
func MailTemplateFkMailHeader(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		MailTemplate.TemHeaderID: goqu.I(MailHeader.HeaID),
	})
	return goqu.On(exps...)
}

// MailTemplateFkMailSender returns the join expression for the foreign key from MailTemplate to MailSender.
func MailTemplateFkMailSender(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		MailTemplate.TemSenderID: goqu.I(MailSender.SenID),
	})
	return goqu.On(exps...)
}

// MailTemplateFkMailTemplateType returns the join expression for the foreign key from MailTemplate to MailTemplateType.
func MailTemplateFkMailTemplateType(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		MailTemplate.TemTypeID: goqu.I(MailTemplateType.TetID),
	})
	return goqu.On(exps...)
}
