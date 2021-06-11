// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dalhelper

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

// T database table names.
var T = struct {
	CoreAccountAccess  string
	CoreAction         string
	CoreActionRoute    string
	CoreModule         string
	CoreNavElement     string
	CoreNavElementRole string
	CorePassword       string
	CorePrivilege      string
	CoreRole           string
	CoreUser           string
	CoreUserRole       string
	MailBRecord        string
	MailFooter         string
	MailHeader         string
	MailSender         string
	MailTemplate       string
	MailTemplateType   string
}{
	CoreAccountAccess:  "core_account_access",
	CoreAction:         "core_action",
	CoreActionRoute:    "core_action_route",
	CoreModule:         "core_module",
	CoreNavElement:     "core_nav_element",
	CoreNavElementRole: "core_nav_element_role",
	CorePassword:       "core_password",
	CorePrivilege:      "core_privilege",
	CoreRole:           "core_role",
	CoreUser:           "core_user",
	CoreUserRole:       "core_user_role",
	MailBRecord:        "mail_b_record",
	MailFooter:         "mail_footer",
	MailHeader:         "mail_header",
	MailSender:         "mail_sender",
	MailTemplate:       "mail_template",
	MailTemplateType:   "mail_template_type",
}

// V database view names.
var V = struct {
	CoreVPrivilegeRole string
}{
	CoreVPrivilegeRole: "core_v_privilege_role",
}

// CoreAccountAccessFkCoreUser returns the join expression for the foreign key from CoreAccountAccess to CoreUser.
func CoreAccountAccessFkCoreUser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CoreAccountAccess.AcaUserID: goqu.I(CoreUser.UseID),
	})
	return goqu.On(exps...)
}

// CoreActionFkCoreModule returns the join expression for the foreign key from CoreAction to CoreModule.
func CoreActionFkCoreModule(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CoreAction.ActModuleID: goqu.I(CoreModule.ModID),
	})
	return goqu.On(exps...)
}

// CoreActionRouteFkCoreAction returns the join expression for the foreign key from CoreActionRoute to CoreAction.
func CoreActionRouteFkCoreAction(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CoreActionRoute.AcrActionID: goqu.I(CoreAction.ActActionID),
		CoreActionRoute.AcrModuleID: goqu.I(CoreAction.ActModuleID),
	})
	return goqu.On(exps...)
}

// CoreModuleFkCoreModulePadre returns the join expression for the foreign key from CoreModule to CoreModule.
func CoreModuleFkCoreModulePadre(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CoreModule.ModParentID: goqu.I(CoreModule.ModID),
	})
	return goqu.On(exps...)
}

// CoreNavElementFkCoreNavElementParent returns the join expression for the foreign key from CoreNavElement to CoreNavElement.
func CoreNavElementFkCoreNavElementParent(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CoreNavElement.NaeParentID: goqu.I(CoreNavElement.NaeID),
	})
	return goqu.On(exps...)
}

// CoreNavElementRoleFkCoreNavElement returns the join expression for the foreign key from CoreNavElementRole to CoreNavElement.
func CoreNavElementRoleFkCoreNavElement(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CoreNavElementRole.NerNavElementID: goqu.I(CoreNavElement.NaeID),
	})
	return goqu.On(exps...)
}

// CoreNavElementRoleFkCoreRole returns the join expression for the foreign key from CoreNavElementRole to CoreRole.
func CoreNavElementRoleFkCoreRole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CoreNavElementRole.NerRoleID: goqu.I(CoreRole.RolID),
	})
	return goqu.On(exps...)
}

// CorePasswordFkCoreUser returns the join expression for the foreign key from CorePassword to CoreUser.
func CorePasswordFkCoreUser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CorePassword.PasUserID: goqu.I(CoreUser.UseID),
	})
	return goqu.On(exps...)
}

// CorePrivilegeFkCoreAction returns the join expression for the foreign key from CorePrivilege to CoreAction.
func CorePrivilegeFkCoreAction(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CorePrivilege.PriActionID: goqu.I(CoreAction.ActActionID),
		CorePrivilege.PriModuleID: goqu.I(CoreAction.ActModuleID),
	})
	return goqu.On(exps...)
}

// CorePrivilegeFkCoreRole returns the join expression for the foreign key from CorePrivilege to CoreRole.
func CorePrivilegeFkCoreRole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CorePrivilege.PriRoleID: goqu.I(CoreRole.RolID),
	})
	return goqu.On(exps...)
}

// CoreUserRoleFkCoreRole returns the join expression for the foreign key from CoreUserRole to CoreRole.
func CoreUserRoleFkCoreRole(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CoreUserRole.UsrRoleID: goqu.I(CoreRole.RolID),
	})
	return goqu.On(exps...)
}

// CoreUserRoleFkCoreUser returns the join expression for the foreign key from CoreUserRole to CoreUser.
func CoreUserRoleFkCoreUser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		CoreUserRole.UsrUserID: goqu.I(CoreUser.UseID),
	})
	return goqu.On(exps...)
}

// MailBRecordFkCoreUser returns the join expression for the foreign key from MailBRecord to CoreUser.
func MailBRecordFkCoreUser(exps ...exp.Expression) exp.JoinCondition {
	exps = append(exps, goqu.Ex{
		MailBRecord.RecSenderUserID: goqu.I(CoreUser.UseID),
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
