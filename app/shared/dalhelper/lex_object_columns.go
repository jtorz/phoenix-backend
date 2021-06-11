// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dalhelper

// TableCoreAccountAccess column names for table core_account_access.
type TableCoreAccountAccess struct {
	AcaID             string `database:"-,datatype=text"`
	AcaType           string `database:"-,datatype=text"`
	AcaUserID         string `database:"-,datatype=uuid"`
	AcaExpirationDate string `database:"-,datatype=timestamp with time zone"`
	AcaCreatedAt      string `database:"-,datatype=timestamp with time zone"`
	AcaUpdatedAt      string `database:"-,datatype=timestamp with time zone"`
	AcaStatus         string `database:"-,datatype=smallint"`
}

var CoreAccountAccess = TableCoreAccountAccess{
	AcaID:             "aca_id",
	AcaType:           "aca_type",
	AcaUserID:         "aca_user_id",
	AcaExpirationDate: "aca_expiration_date",
	AcaCreatedAt:      "aca_created_at",
	AcaUpdatedAt:      "aca_updated_at",
	AcaStatus:         "aca_status",
}

// TableCoreAction column names for table core_action.
type TableCoreAction struct {
	ActModuleID    string `database:"-,datatype=text"`
	ActActionID    string `database:"-,datatype=text"`
	ActName        string `database:"-,datatype=text"`
	ActDescription string `database:"-,datatype=text"`
	ActOrder       string `database:"-,datatype=integer"`
	ActCreatedAt   string `database:"-,datatype=timestamp with time zone"`
	ActUpdatedAt   string `database:"-,datatype=timestamp with time zone"`
	ActStatus      string `database:"-,datatype=smallint"`
}

var CoreAction = TableCoreAction{
	ActModuleID:    "act_module_id",
	ActActionID:    "act_action_id",
	ActName:        "act_name",
	ActDescription: "act_description",
	ActOrder:       "act_order",
	ActCreatedAt:   "act_created_at",
	ActUpdatedAt:   "act_updated_at",
	ActStatus:      "act_status",
}

// TableCoreActionRoute column names for table core_action_route.
type TableCoreActionRoute struct {
	AcrModuleID string `database:"-,datatype=text"`
	AcrActionID string `database:"-,datatype=text"`
	AcrMethod   string `database:"-,datatype=text"`
	AcrRoute    string `database:"-,datatype=text"`
}

var CoreActionRoute = TableCoreActionRoute{
	AcrModuleID: "acr_module_id",
	AcrActionID: "acr_action_id",
	AcrMethod:   "acr_method",
	AcrRoute:    "acr_route",
}

// TableCoreModule column names for table core_module.
type TableCoreModule struct {
	ModID          string `database:"-,datatype=text"`
	ModName        string `database:"-,datatype=text"`
	ModDescription string `database:"-,datatype=text"`
	ModOrder       string `database:"-,datatype=integer"`
	ModParentID    string `database:"N,datatype=text"`
	ModCreatedAt   string `database:"-,datatype=timestamp with time zone"`
	ModUpdatedAt   string `database:"-,datatype=timestamp with time zone"`
	ModStatus      string `database:"-,datatype=smallint"`
}

var CoreModule = TableCoreModule{
	ModID:          "mod_id",
	ModName:        "mod_name",
	ModDescription: "mod_description",
	ModOrder:       "mod_order",
	ModParentID:    "mod_parent_id",
	ModCreatedAt:   "mod_created_at",
	ModUpdatedAt:   "mod_updated_at",
	ModStatus:      "mod_status",
}

// TableCoreNavElement column names for table core_nav_element.
type TableCoreNavElement struct {
	NaeID          string `database:"-,datatype=text"`
	NaeName        string `database:"-,datatype=text"`
	NaeDescription string `database:"-,datatype=text"`
	NaeIcon        string `database:"-,datatype=text"`
	NaeOrder       string `database:"-,datatype=integer"`
	NaeURL         string `database:"-,datatype=text"`
	NaeParentID    string `database:"N,datatype=text"`
	NaeCreatedAt   string `database:"-,datatype=timestamp with time zone"`
	NaeUpdatedAt   string `database:"-,datatype=timestamp with time zone"`
	NaeStatus      string `database:"-,datatype=smallint"`
}

var CoreNavElement = TableCoreNavElement{
	NaeID:          "nae_id",
	NaeName:        "nae_name",
	NaeDescription: "nae_description",
	NaeIcon:        "nae_icon",
	NaeOrder:       "nae_order",
	NaeURL:         "nae_url",
	NaeParentID:    "nae_parent_id",
	NaeCreatedAt:   "nae_created_at",
	NaeUpdatedAt:   "nae_updated_at",
	NaeStatus:      "nae_status",
}

// TableCoreNavElementRole column names for table core_nav_element_role.
type TableCoreNavElementRole struct {
	NerNavElementID string `database:"-,datatype=text"`
	NerRoleID       string `database:"-,datatype=text"`
}

var CoreNavElementRole = TableCoreNavElementRole{
	NerNavElementID: "ner_nav_element_id",
	NerRoleID:       "ner_role_id",
}

// TableCorePassword column names for table core_password.
type TableCorePassword struct {
	PasID               string `database:"-,datatype=bigint"`
	PasData             string `database:"-,datatype=json"`
	PasType             string `database:"-,datatype=text"`
	PasUserID           string `database:"-,datatype=uuid"`
	PasInvalidationDate string `database:"N,datatype=timestamp with time zone"`
	PasCreatedAt        string `database:"-,datatype=timestamp with time zone"`
	PasUpdatedAt        string `database:"-,datatype=timestamp with time zone"`
	PasStatus           string `database:"-,datatype=smallint"`
}

var CorePassword = TableCorePassword{
	PasID:               "pas_id",
	PasData:             "pas_data",
	PasType:             "pas_type",
	PasUserID:           "pas_user_id",
	PasInvalidationDate: "pas_invalidation_date",
	PasCreatedAt:        "pas_created_at",
	PasUpdatedAt:        "pas_updated_at",
	PasStatus:           "pas_status",
}

// TableCorePrivilege column names for table core_privilege.
type TableCorePrivilege struct {
	PriRoleID   string `database:"-,datatype=text"`
	PriModuleID string `database:"-,datatype=text"`
	PriActionID string `database:"-,datatype=text"`
}

var CorePrivilege = TableCorePrivilege{
	PriRoleID:   "pri_role_id",
	PriModuleID: "pri_module_id",
	PriActionID: "pri_action_id",
}

// TableCoreRole column names for table core_role.
type TableCoreRole struct {
	RolID          string `database:"-,datatype=text"`
	RolName        string `database:"-,datatype=text"`
	RolDescription string `database:"-,datatype=text"`
	RolIcon        string `database:"-,datatype=text"`
	RolCreatedAt   string `database:"-,datatype=timestamp with time zone"`
	RolUpdatedAt   string `database:"-,datatype=timestamp with time zone"`
	RolStatus      string `database:"-,datatype=smallint"`
}

var CoreRole = TableCoreRole{
	RolID:          "rol_id",
	RolName:        "rol_name",
	RolDescription: "rol_description",
	RolIcon:        "rol_icon",
	RolCreatedAt:   "rol_created_at",
	RolUpdatedAt:   "rol_updated_at",
	RolStatus:      "rol_status",
}

// TableCoreUser column names for table core_user.
type TableCoreUser struct {
	UseID         string `database:"-,datatype=uuid"`
	UseName       string `database:"-,datatype=text"`
	UseMiddleName string `database:"-,datatype=text"`
	UseLastName   string `database:"-,datatype=text"`
	UseEmail      string `database:"-,datatype=text"`
	UseUsername   string `database:"-,datatype=text"`
	UseCreatedAt  string `database:"-,datatype=timestamp with time zone"`
	UseUpdatedAt  string `database:"-,datatype=timestamp with time zone"`
	UseStatus     string `database:"-,datatype=smallint"`
}

var CoreUser = TableCoreUser{
	UseID:         "use_id",
	UseName:       "use_name",
	UseMiddleName: "use_middle_name",
	UseLastName:   "use_last_name",
	UseEmail:      "use_email",
	UseUsername:   "use_username",
	UseCreatedAt:  "use_created_at",
	UseUpdatedAt:  "use_updated_at",
	UseStatus:     "use_status",
}

// TableCoreUserRole column names for table core_user_role.
type TableCoreUserRole struct {
	UsrUserID string `database:"-,datatype=uuid"`
	UsrRoleID string `database:"-,datatype=text"`
}

var CoreUserRole = TableCoreUserRole{
	UsrUserID: "usr_user_id",
	UsrRoleID: "usr_role_id",
}

// TableMailBRecord column names for table mail_b_record.
type TableMailBRecord struct {
	RecID           string `database:"-,datatype=bigint"`
	RecTypeID       string `database:"N,datatype=text"`
	RecEmail        string `database:"-,datatype=text"`
	RecSenderUserID string `database:"N,datatype=uuid"`
	RecError        string `database:"N,datatype=text"`
	RecTo           string `database:"-,datatype=text"`
	RecCc           string `database:"N,datatype=text"`
	RecBcc          string `database:"N,datatype=text"`
	RecSubject      string `database:"-,datatype=text"`
	RecMime         string `database:"N,datatype=text"`
	RecFrom         string `database:"-,datatype=text"`
	RecCreatedAt    string `database:"-,datatype=timestamp with time zone"`
	RecUpdatedAt    string `database:"-,datatype=timestamp with time zone"`
	RecStatus       string `database:"-,datatype=smallint"`
}

var MailBRecord = TableMailBRecord{
	RecID:           "rec_id",
	RecTypeID:       "rec_type_id",
	RecEmail:        "rec_email",
	RecSenderUserID: "rec_sender_user_id",
	RecError:        "rec_error",
	RecTo:           "rec_to",
	RecCc:           "rec_cc",
	RecBcc:          "rec_bcc",
	RecSubject:      "rec_subject",
	RecMime:         "rec_mime",
	RecFrom:         "rec_from",
	RecCreatedAt:    "rec_created_at",
	RecUpdatedAt:    "rec_updated_at",
	RecStatus:       "rec_status",
}

// TableMailFooter column names for table mail_footer.
type TableMailFooter struct {
	FooID        string `database:"-,datatype=integer"`
	FooName      string `database:"-,datatype=text"`
	FooFooter    string `database:"-,datatype=text"`
	FooCreatedAt string `database:"-,datatype=timestamp with time zone"`
	FooUpdatedAt string `database:"-,datatype=timestamp with time zone"`
	FooStatus    string `database:"-,datatype=smallint"`
}

var MailFooter = TableMailFooter{
	FooID:        "foo_id",
	FooName:      "foo_name",
	FooFooter:    "foo_footer",
	FooCreatedAt: "foo_created_at",
	FooUpdatedAt: "foo_updated_at",
	FooStatus:    "foo_status",
}

// TableMailHeader column names for table mail_header.
type TableMailHeader struct {
	HeaID        string `database:"-,datatype=integer"`
	HeaName      string `database:"-,datatype=text"`
	HeaHeader    string `database:"-,datatype=text"`
	HeaCreatedAt string `database:"-,datatype=timestamp with time zone"`
	HeaUpdatedAt string `database:"-,datatype=timestamp with time zone"`
	HeaStatus    string `database:"-,datatype=smallint"`
}

var MailHeader = TableMailHeader{
	HeaID:        "hea_id",
	HeaName:      "hea_name",
	HeaHeader:    "hea_header",
	HeaCreatedAt: "hea_created_at",
	HeaUpdatedAt: "hea_updated_at",
	HeaStatus:    "hea_status",
}

// TableMailSender column names for table mail_sender.
type TableMailSender struct {
	SenID          string `database:"-,datatype=text"`
	SenName        string `database:"-,datatype=text"`
	SenDescription string `database:"-,datatype=text"`
	SenHost        string `database:"-,datatype=text"`
	SenPort        string `database:"-,datatype=integer"`
	SenUser        string `database:"-,datatype=text"`
	SenPassword    string `database:"-,datatype=text"`
	SenFrom        string `database:"-,datatype=text"`
	SenCreatedAt   string `database:"-,datatype=timestamp with time zone"`
	SenUpdatedAt   string `database:"-,datatype=timestamp with time zone"`
	SenStatus      string `database:"-,datatype=smallint"`
}

var MailSender = TableMailSender{
	SenID:          "sen_id",
	SenName:        "sen_name",
	SenDescription: "sen_description",
	SenHost:        "sen_host",
	SenPort:        "sen_port",
	SenUser:        "sen_user",
	SenPassword:    "sen_password",
	SenFrom:        "sen_from",
	SenCreatedAt:   "sen_created_at",
	SenUpdatedAt:   "sen_updated_at",
	SenStatus:      "sen_status",
}

// TableMailTemplate column names for table mail_template.
type TableMailTemplate struct {
	TemTypeID     string `database:"-,datatype=text"`
	TemTemplateID string `database:"-,datatype=integer"`
	TemFrom       string `database:"-,datatype=text"`
	TemSubject    string `database:"-,datatype=text"`
	TemTemplate   string `database:"-,datatype=text"`
	TemHeaderID   string `database:"-,datatype=integer"`
	TemFooterID   string `database:"-,datatype=integer"`
	TemSenderID   string `database:"-,datatype=text"`
	TemCreatedAt  string `database:"-,datatype=timestamp with time zone"`
	TemUpdatedAt  string `database:"-,datatype=timestamp with time zone"`
	TemStatus     string `database:"-,datatype=smallint"`
}

var MailTemplate = TableMailTemplate{
	TemTypeID:     "tem_type_id",
	TemTemplateID: "tem_template_id",
	TemFrom:       "tem_from",
	TemSubject:    "tem_subject",
	TemTemplate:   "tem_template",
	TemHeaderID:   "tem_header_id",
	TemFooterID:   "tem_footer_id",
	TemSenderID:   "tem_sender_id",
	TemCreatedAt:  "tem_created_at",
	TemUpdatedAt:  "tem_updated_at",
	TemStatus:     "tem_status",
}

// TableMailTemplateType column names for table mail_template_type.
type TableMailTemplateType struct {
	TetID        string `database:"-,datatype=text"`
	TetName      string `database:"-,datatype=text"`
	TetTags      string `database:"-,datatype=json"`
	TetCreatedAt string `database:"-,datatype=timestamp with time zone"`
	TetUpdatedAt string `database:"-,datatype=timestamp with time zone"`
	TetStatus    string `database:"-,datatype=smallint"`
}

var MailTemplateType = TableMailTemplateType{
	TetID:        "tet_id",
	TetName:      "tet_name",
	TetTags:      "tet_tags",
	TetCreatedAt: "tet_created_at",
	TetUpdatedAt: "tet_updated_at",
	TetStatus:    "tet_status",
}

// ViewCoreVPrivilegeRole column names for view core_v_privilege_role.
type ViewCoreVPrivilegeRole struct {
	PrrRoleID   string `database:"N,datatype=text"`
	PrrModuleID string `database:"N,datatype=text"`
	PrrActionID string `database:"N,datatype=text"`
	PrrMethod   string `database:"N,datatype=text"`
	PrrRoute    string `database:"N,datatype=text"`
}

var CoreVPrivilegeRole = ViewCoreVPrivilegeRole{
	PrrRoleID:   "prr_role_id",
	PrrModuleID: "prr_module_id",
	PrrActionID: "prr_action_id",
	PrrMethod:   "prr_method",
	PrrRoute:    "prr_route",
}
