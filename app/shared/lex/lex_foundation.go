package lex

var (
	// Fndccomponent table column names.
	Fndccomponent = struct {
		ComID          string
		ComName        string
		ComDescription string
		ComCreatedAt   string
		ComUpdatedAt   string
		ComStatus      string
	}{
		ComID:          "com_id",
		ComName:        "com_name",
		ComDescription: "com_description",
		ComCreatedAt:   "com_created_at",
		ComUpdatedAt:   "com_updated_at",
		ComStatus:      "com_status",
	}

	// Fndtmodule table column names.
	Fndtmodule = struct {
		ModID          string
		ModName        string
		ModDescription string
		ModCreatedAt   string
		ModUpdatedAt   string
		ModStatus      string
	}{
		ModID:          "mod_id",
		ModName:        "mod_name",
		ModDescription: "mod_description",
		ModCreatedAt:   "mod_created_at",
		ModUpdatedAt:   "mod_updated_at",
		ModStatus:      "mod_status",
	}

	// Fndtaction table column names.
	Fndtaction = struct {
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

	// Fndtrole table column names.
	Fndtrole = struct {
		RolID          string
		RolName        string
		RolDescription string
		RolCreatedAt   string
		RolUpdatedAt   string
		RolStatus      string
	}{
		RolID:          "rol_id",
		RolName:        "rol_name",
		RolDescription: "rol_description",
		RolCreatedAt:   "rol_created_at",
		RolUpdatedAt:   "rol_updated_at",
		RolStatus:      "rol_status",
	}

	// Fndtprivilege table column names.
	Fndtprivilege = struct {
		PriRoleID   string
		PriModuleID string
		PriActionID string
	}{
		PriRoleID:   "pri_role_id",
		PriModuleID: "act_module_id",
		PriActionID: "act_action_id",
	}

	// Fndtuser table column names.
	Fndtuser = struct {
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

	// Fndtpassword table column names.
	Fndtpassword = struct {
		PasID               string
		PasData             string
		PasType             string
		PasInvalidationDate string
		PasUserID           string
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

	// FndtuserRole table column names.
	FndtuserRole = struct {
		UsrUserID string
		UsrRoleID string
	}{
		UsrUserID: "usr_user_id",
		UsrRoleID: "usr_role_id",
	}

	// Fndtnavigator table column names.
	Fndtnavigator = struct {
		NavID          string
		NavName        string
		NavDescription string
		NavIcon        string
		NavOrder       string
		NavUrl         string
		NavCreatedAt   string
		NavUpdatedAt   string
		NavStatus      string
	}{
		NavID:          "nav_id",
		NavName:        "nav_name",
		NavDescription: "nav_description",
		NavIcon:        "nav_icon",
		NavOrder:       "nav_order",
		NavUrl:         "nav_url",
		NavCreatedAt:   "nav_created_at",
		NavUpdatedAt:   "nav_updated_at",
		NavStatus:      "nav_status",
	}

	// FndtroleNavigator table column names.
	FndtroleNavigator = struct {
		RonRoleID      string
		RonNavigatorID string
	}{
		RonRoleID:      "ron_role_id",
		RonNavigatorID: "ron_navigator_id",
	}

	// FndtaccessAccount table column names.
	FndtaccessAccount = struct {
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

	// FndvprivilegeRole table column names.
	FndvprivilegeRole = struct {
		PrrRoleID   string
		PrrModuleID string
		PrrActionID string
		PrrRoute    string
		PrrMethod   string
	}{
		PrrRoleID:   "prr_role_id",
		PrrModuleID: "prr_module_id",
		PrrActionID: "prr_action_id",
		PrrRoute:    "prr_route",
		PrrMethod:   "prr_method",
	}
)
