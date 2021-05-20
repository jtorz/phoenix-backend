package lex

var (
	// FndModule column names for object fnd_module.
	FndModule = struct {
		ModID          string
		ModName        string
		ModDescription string
		ModOrder       string
		ModParentID    string
		ModCreatedAt   string
		ModUpdatedAt   string
		ModStatus      string
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

	// FndAction column names for object fnd_action.
	FndAction = struct {
		ActModuleID    string
		ActActionID    string
		ActName        string
		ActDescription string
		ActOrder       string
		ActRoute       string
		ActMethod      string
		ActCreatedAt   string
		ActUpdatedAt   string
		ActStatus      string
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

	// FndRole column names for object fnd_role.
	FndRole = struct {
		RolID          string
		RolName        string
		RolDescription string
		RolIcon        string
		RolCreatedAt   string
		RolUpdatedAt   string
		RolStatus      string
	}{
		RolID:          "rol_id",
		RolName:        "rol_name",
		RolDescription: "rol_description",
		RolIcon:        "rol_icon",
		RolCreatedAt:   "rol_created_at",
		RolUpdatedAt:   "rol_updated_at",
		RolStatus:      "rol_status",
	}

	// FndPrivilege column names for object fnd_privilege.
	FndPrivilege = struct {
		PriRoleID   string
		PriModuleID string
		PriActionID string
	}{
		PriRoleID:   "pri_role_id",
		PriModuleID: "pri_module_id",
		PriActionID: "pri_action_id",
	}

	// FndUser column names for object fnd_user.
	FndUser = struct {
		UseID         string
		UseName       string
		UseMiddleName string
		UseLastName   string
		UseEmail      string
		UseUsername   string
		UseCreatedAt  string
		UseUpdatedAt  string
		UseStatus     string
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

	// FndPassword column names for object fnd_password.
	FndPassword = struct {
		PasID               string
		PasData             string
		PasType             string
		PasUserID           string
		PasInvalidationDate string
		PasCreatedAt        string
		PasUpdatedAt        string
		PasStatus           string
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

	// FndUserRole column names for object fnd_user_role.
	FndUserRole = struct {
		UsrUserID string
		UsrRoleID string
	}{
		UsrUserID: "usr_user_id",
		UsrRoleID: "usr_role_id",
	}

	// FndRoleNavigator column names for object fnd_role_navigator.
	FndRoleNavigator = struct {
		RonRoleID      string
		RonNavigatorID string
	}{
		RonRoleID:      "ron_role_id",
		RonNavigatorID: "ron_navigator_id",
	}

	// FndNavigator column names for object fnd_navigator.
	FndNavigator = struct {
		NavID          string
		NavName        string
		NavDescription string
		NavIcon        string
		NavOrder       string
		NavURL         string
		NavCreatedAt   string
		NavUpdatedAt   string
		NavStatus      string
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

	// FndAccountAccess column names for object fnd_account_access.
	FndAccountAccess = struct {
		AcaID             string
		AcaType           string
		AcaUserID         string
		AcaExpirationDate string
		AcaCreatedAt      string
		AcaUpdatedAt      string
		AcaStatus         string
	}{
		AcaID:             "aca_id",
		AcaType:           "aca_type",
		AcaUserID:         "aca_user_id",
		AcaExpirationDate: "aca_expiration_date",
		AcaCreatedAt:      "aca_created_at",
		AcaUpdatedAt:      "aca_updated_at",
		AcaStatus:         "aca_status",
	}
)
