// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lex

// TableFndAccountAccess column names for table fnd_account_access.
type TableFndAccountAccess struct {
	AcaID             string `database:"-,datatype=text"`
	AcaType           string `database:"-,datatype=text"`
	AcaUserID         string `database:"-,datatype=uuid"`
	AcaExpirationDate string `database:"-,datatype=timestamp with time zone"`
	AcaCreatedAt      string `database:"-,datatype=timestamp with time zone"`
	AcaUpdatedAt      string `database:"-,datatype=timestamp with time zone"`
	AcaStatus         string `database:"-,datatype=smallint"`
}

var FndAccountAccess = TableFndAccountAccess{
	AcaID:             "aca_id",
	AcaType:           "aca_type",
	AcaUserID:         "aca_user_id",
	AcaExpirationDate: "aca_expiration_date",
	AcaCreatedAt:      "aca_created_at",
	AcaUpdatedAt:      "aca_updated_at",
	AcaStatus:         "aca_status",
}

// TableFndAction column names for table fnd_action.
type TableFndAction struct {
	ActModuleID    string `database:"-,datatype=text"`
	ActActionID    string `database:"-,datatype=text"`
	ActName        string `database:"-,datatype=text"`
	ActDescription string `database:"-,datatype=text"`
	ActOrder       string `database:"-,datatype=integer"`
	ActRoute       string `database:"-,datatype=text"`
	ActMethod      string `database:"-,datatype=text"`
	ActCreatedAt   string `database:"-,datatype=timestamp with time zone"`
	ActUpdatedAt   string `database:"-,datatype=timestamp with time zone"`
	ActStatus      string `database:"-,datatype=smallint"`
}

var FndAction = TableFndAction{
	ActModuleID:    "act_module_id",
	ActActionID:    "act_action_id",
	ActName:        "act_name",
	ActDescription: "act_description",
	ActOrder:       "act_order",
	ActRoute:       "act_route",
	ActMethod:      "act_method",
	ActCreatedAt:   "act_created_at",
	ActUpdatedAt:   "act_updated_at",
	ActStatus:      "act_status",
}

// TableFndModule column names for table fnd_module.
type TableFndModule struct {
	ModID          string `database:"-,datatype=text"`
	ModName        string `database:"-,datatype=text"`
	ModDescription string `database:"-,datatype=text"`
	ModOrder       string `database:"-,datatype=integer"`
	ModParentID    string `database:"N,datatype=text"`
	ModCreatedAt   string `database:"-,datatype=timestamp with time zone"`
	ModUpdatedAt   string `database:"-,datatype=timestamp with time zone"`
	ModStatus      string `database:"-,datatype=smallint"`
}

var FndModule = TableFndModule{
	ModID:          "mod_id",
	ModName:        "mod_name",
	ModDescription: "mod_description",
	ModOrder:       "mod_order",
	ModParentID:    "mod_parent_id",
	ModCreatedAt:   "mod_created_at",
	ModUpdatedAt:   "mod_updated_at",
	ModStatus:      "mod_status",
}

// TableFndNavigator column names for table fnd_navigator.
type TableFndNavigator struct {
	NavID          string `database:"-,datatype=text"`
	NavName        string `database:"-,datatype=text"`
	NavDescription string `database:"-,datatype=text"`
	NavIcon        string `database:"-,datatype=text"`
	NavOrder       string `database:"-,datatype=text"`
	NavURL         string `database:"-,datatype=text"`
	NavParentID    string `database:"N,datatype=text"`
	NavCreatedAt   string `database:"-,datatype=timestamp with time zone"`
	NavUpdatedAt   string `database:"-,datatype=timestamp with time zone"`
	NavStatus      string `database:"-,datatype=smallint"`
}

var FndNavigator = TableFndNavigator{
	NavID:          "nav_id",
	NavName:        "nav_name",
	NavDescription: "nav_description",
	NavIcon:        "nav_icon",
	NavOrder:       "nav_order",
	NavURL:         "nav_url",
	NavParentID:    "nav_parent_id",
	NavCreatedAt:   "nav_created_at",
	NavUpdatedAt:   "nav_updated_at",
	NavStatus:      "nav_status",
}

// TableFndPassword column names for table fnd_password.
type TableFndPassword struct {
	PasID               string `database:"-,datatype=bigint"`
	PasData             string `database:"-,datatype=json"`
	PasType             string `database:"-,datatype=text"`
	PasUserID           string `database:"-,datatype=uuid"`
	PasInvalidationDate string `database:"N,datatype=timestamp with time zone"`
	PasCreatedAt        string `database:"-,datatype=timestamp with time zone"`
	PasUpdatedAt        string `database:"-,datatype=timestamp with time zone"`
	PasStatus           string `database:"-,datatype=smallint"`
}

var FndPassword = TableFndPassword{
	PasID:               "pas_id",
	PasData:             "pas_data",
	PasType:             "pas_type",
	PasUserID:           "pas_user_id",
	PasInvalidationDate: "pas_invalidation_date",
	PasCreatedAt:        "pas_created_at",
	PasUpdatedAt:        "pas_updated_at",
	PasStatus:           "pas_status",
}

// TableFndPrivilege column names for table fnd_privilege.
type TableFndPrivilege struct {
	PriRoleID   string `database:"-,datatype=text"`
	PriModuleID string `database:"-,datatype=text"`
	PriActionID string `database:"-,datatype=text"`
}

var FndPrivilege = TableFndPrivilege{
	PriRoleID:   "pri_role_id",
	PriModuleID: "pri_module_id",
	PriActionID: "pri_action_id",
}

// TableFndRole column names for table fnd_role.
type TableFndRole struct {
	RolID          string `database:"-,datatype=text"`
	RolName        string `database:"-,datatype=text"`
	RolDescription string `database:"-,datatype=text"`
	RolIcon        string `database:"-,datatype=text"`
	RolCreatedAt   string `database:"-,datatype=timestamp with time zone"`
	RolUpdatedAt   string `database:"-,datatype=timestamp with time zone"`
	RolStatus      string `database:"-,datatype=smallint"`
}

var FndRole = TableFndRole{
	RolID:          "rol_id",
	RolName:        "rol_name",
	RolDescription: "rol_description",
	RolIcon:        "rol_icon",
	RolCreatedAt:   "rol_created_at",
	RolUpdatedAt:   "rol_updated_at",
	RolStatus:      "rol_status",
}

// TableFndRoleNavigator column names for table fnd_role_navigator.
type TableFndRoleNavigator struct {
	RonRoleID      string `database:"-,datatype=text"`
	RonNavigatorID string `database:"-,datatype=text"`
}

var FndRoleNavigator = TableFndRoleNavigator{
	RonRoleID:      "ron_role_id",
	RonNavigatorID: "ron_navigator_id",
}

// TableFndUser column names for table fnd_user.
type TableFndUser struct {
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

var FndUser = TableFndUser{
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

// TableFndUserRole column names for table fnd_user_role.
type TableFndUserRole struct {
	UsrUserID string `database:"-,datatype=uuid"`
	UsrRoleID string `database:"-,datatype=text"`
}

var FndUserRole = TableFndUserRole{
	UsrUserID: "usr_user_id",
	UsrRoleID: "usr_role_id",
}

// ViewFndVPrivilegeRole column names for view fnd_v_privilege_role.
type ViewFndVPrivilegeRole struct {
	PrrRoleID   string `database:"N,datatype=text"`
	PrrModuleID string `database:"N,datatype=text"`
	PrrActionID string `database:"N,datatype=text"`
	PrrRoute    string `database:"N,datatype=text"`
	PrrMethod   string `database:"N,datatype=text"`
}

var FndVPrivilegeRole = ViewFndVPrivilegeRole{
	PrrRoleID:   "prr_role_id",
	PrrModuleID: "prr_module_id",
	PrrActionID: "prr_action_id",
	PrrRoute:    "prr_route",
	PrrMethod:   "prr_method",
}
