// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package lex

var (
	// FndModule column names for table fnd_module.
	FndModule = struct {
		ModID          string `database:"-,text"`
		ModName        string `database:"-,text"`
		ModDescription string `database:"-,text"`
		ModOrder       string `database:"-,integer"`
		ModParentID    string `database:"N,text"`
		ModCreatedAt   string `database:"-,timestamp with time zone"`
		ModUpdatedAt   string `database:"-,timestamp with time zone"`
		ModStatus      string `database:"-,smallint"`
	}{
		ModID:          "mod_id",
		ModName:        "mod_name",
		ModDescription: "mod_description",
		ModOrder:       "mod_order",
		ModParentID:    "mod_parent_id",
		ModCreatedAt:   "mod_created_at",
		ModUpdatedAt:   "mod_updated_at",
		ModStatus:      "mod_status",
	}

	// FndAction column names for table fnd_action.
	FndAction = struct {
		ActModuleID    string `database:"-,text"`
		ActActionID    string `database:"-,text"`
		ActName        string `database:"-,text"`
		ActDescription string `database:"-,text"`
		ActOrder       string `database:"-,integer"`
		ActRoute       string `database:"-,text"`
		ActMethod      string `database:"-,text"`
		ActCreatedAt   string `database:"-,timestamp with time zone"`
		ActUpdatedAt   string `database:"-,timestamp with time zone"`
		ActStatus      string `database:"-,smallint"`
	}{
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

	// FndRole column names for table fnd_role.
	FndRole = struct {
		RolID          string `database:"-,text"`
		RolName        string `database:"-,text"`
		RolDescription string `database:"-,text"`
		RolIcon        string `database:"-,text"`
		RolCreatedAt   string `database:"-,timestamp with time zone"`
		RolUpdatedAt   string `database:"-,timestamp with time zone"`
		RolStatus      string `database:"-,smallint"`
	}{
		RolID:          "rol_id",
		RolName:        "rol_name",
		RolDescription: "rol_description",
		RolIcon:        "rol_icon",
		RolCreatedAt:   "rol_created_at",
		RolUpdatedAt:   "rol_updated_at",
		RolStatus:      "rol_status",
	}

	// FndPrivilege column names for table fnd_privilege.
	FndPrivilege = struct {
		PriRoleID   string `database:"-,text"`
		PriModuleID string `database:"-,text"`
		PriActionID string `database:"-,text"`
	}{
		PriRoleID:   "pri_role_id",
		PriModuleID: "pri_module_id",
		PriActionID: "pri_action_id",
	}

	// FndUser column names for table fnd_user.
	FndUser = struct {
		UseID         string `database:"-,uuid"`
		UseName       string `database:"-,text"`
		UseMiddleName string `database:"-,text"`
		UseLastName   string `database:"-,text"`
		UseEmail      string `database:"-,text"`
		UseUsername   string `database:"-,text"`
		UseCreatedAt  string `database:"-,timestamp with time zone"`
		UseUpdatedAt  string `database:"-,timestamp with time zone"`
		UseStatus     string `database:"-,smallint"`
	}{
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

	// FndPassword column names for table fnd_password.
	FndPassword = struct {
		PasID               string `database:"-,bigint"`
		PasData             string `database:"-,json"`
		PasType             string `database:"-,text"`
		PasUserID           string `database:"-,uuid"`
		PasInvalidationDate string `database:"N,timestamp with time zone"`
		PasCreatedAt        string `database:"-,timestamp with time zone"`
		PasUpdatedAt        string `database:"-,timestamp with time zone"`
		PasStatus           string `database:"-,smallint"`
	}{
		PasID:               "pas_id",
		PasData:             "pas_data",
		PasType:             "pas_type",
		PasUserID:           "pas_user_id",
		PasInvalidationDate: "pas_invalidation_date",
		PasCreatedAt:        "pas_created_at",
		PasUpdatedAt:        "pas_updated_at",
		PasStatus:           "pas_status",
	}

	// FndUserRole column names for table fnd_user_role.
	FndUserRole = struct {
		UsrUserID string `database:"-,uuid"`
		UsrRoleID string `database:"-,text"`
	}{
		UsrUserID: "usr_user_id",
		UsrRoleID: "usr_role_id",
	}

	// FndRoleNavigator column names for table fnd_role_navigator.
	FndRoleNavigator = struct {
		RonRoleID      string `database:"-,text"`
		RonNavigatorID string `database:"-,text"`
	}{
		RonRoleID:      "ron_role_id",
		RonNavigatorID: "ron_navigator_id",
	}

	// FndNavigator column names for table fnd_navigator.
	FndNavigator = struct {
		NavID          string `database:"-,text"`
		NavName        string `database:"-,text"`
		NavDescription string `database:"-,text"`
		NavIcon        string `database:"-,text"`
		NavOrder       string `database:"-,text"`
		NavURL         string `database:"-,text"`
		NavCreatedAt   string `database:"-,timestamp with time zone"`
		NavUpdatedAt   string `database:"-,timestamp with time zone"`
		NavStatus      string `database:"-,smallint"`
	}{
		NavID:          "nav_id",
		NavName:        "nav_name",
		NavDescription: "nav_description",
		NavIcon:        "nav_icon",
		NavOrder:       "nav_order",
		NavURL:         "nav_url",
		NavCreatedAt:   "nav_created_at",
		NavUpdatedAt:   "nav_updated_at",
		NavStatus:      "nav_status",
	}

	// FndAccountAccess column names for table fnd_account_access.
	FndAccountAccess = struct {
		AcaID             string `database:"-,text"`
		AcaType           string `database:"-,text"`
		AcaUserID         string `database:"-,uuid"`
		AcaExpirationDate string `database:"-,timestamp with time zone"`
		AcaCreatedAt      string `database:"-,timestamp with time zone"`
		AcaUpdatedAt      string `database:"-,timestamp with time zone"`
		AcaStatus         string `database:"-,smallint"`
	}{
		AcaID:             "aca_id",
		AcaType:           "aca_type",
		AcaUserID:         "aca_user_id",
		AcaExpirationDate: "aca_expiration_date",
		AcaCreatedAt:      "aca_created_at",
		AcaUpdatedAt:      "aca_updated_at",
		AcaStatus:         "aca_status",
	}

	// FndVPrivilegeRole column names for view fnd_v_privilege_role.
	FndVPrivilegeRole = struct {
		PrrRoleID   string `database:"N,text"`
		PrrModuleID string `database:"N,text"`
		PrrActionID string `database:"N,text"`
		PrrRoute    string `database:"N,text"`
		PrrMethod   string `database:"N,text"`
	}{
		PrrRoleID:   "prr_role_id",
		PrrModuleID: "prr_module_id",
		PrrActionID: "prr_action_id",
		PrrRoute:    "prr_route",
		PrrMethod:   "prr_method",
	}
)
