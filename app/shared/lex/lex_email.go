package lex

var (
	Emacsender = struct {
		SenID          string
		SenName        string
		SenDescription string
		SenHost        string
		SenPort        string
		SenUser        string
		SenPassword    string
		SenFrom        string
		SenCreatedAt   string
		SenUpdatedAt   string
		SenStatus      string
	}{
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

	// EmactemplateType table column names.
	EmactemplateType = struct {
		TetID        string
		TetName      string
		TetTags      string
		TetCreatedAt string
		TetUpdatedAt string
		TetStatus    string
	}{
		TetID:        "tet_id",
		TetName:      "tet_name",
		TetTags:      "tet_tags",
		TetCreatedAt: "tet_created_at",
		TetUpdatedAt: "tet_updated_at",
		TetStatus:    "tet_status",
	}

	// Ematheader table column names.
	Ematheader = struct {
		HeaID        string
		HeaName      string
		HeaHeader    string
		HeaCreatorID string
		HeaCreatedAt string
		HeaUpdatedAt string
		HeaStatus    string
	}{
		HeaID:        "hea_id",
		HeaName:      "hea_name",
		HeaHeader:    "hea_header",
		HeaCreatorID: "hea_creator_id",
		HeaCreatedAt: "hea_created_at",
		HeaUpdatedAt: "hea_updated_at",
		HeaStatus:    "hea_status",
	}

	// Ematfooter table column names.
	Ematfooter = struct {
		FooID        string
		FooName      string
		FooFooter    string
		FooCreatorID string
		FooCreatedAt string
		FooUpdatedAt string
		FooStatus    string
	}{
		FooID:        "foo_id",
		FooName:      "foo_name",
		FooFooter:    "foo_footer",
		FooCreatorID: "foo_creator_id",
		FooCreatedAt: "foo_created_at",
		FooUpdatedAt: "foo_updated_at",
		FooStatus:    "foo_status",
	}

	// Emattemplate table column names.
	Emattemplate = struct {
		TemID        string
		TemCodeID    string
		TemFrom      string
		TemSubject   string
		TemTemplate  string
		TemHeaderID  string
		TemFooterID  string
		TemCreatorID string
		TemSenderID  string
		TemCreatedAt string
		TemUpdatedAt string
		TemStatus    string
	}{
		TemID:        "tem_id",
		TemCodeID:    "tem_code_id",
		TemFrom:      "tem_from",
		TemSubject:   "tem_subject",
		TemTemplate:  "tem_template",
		TemHeaderID:  "tem_header_id",
		TemFooterID:  "tem_footer_id",
		TemCreatorID: "tem_creator_id",
		TemSenderID:  "tem_sender_id",
		TemCreatedAt: "tem_created_at",
		TemUpdatedAt: "tem_updated_at",
		TemStatus:    "tem_status",
	}

	// Emabrecord table column names.
	Emabrecord = struct {
		RecID           string
		RecEmail        string
		RecTemplateID   string
		RecSenderUserID string
		RecError        string
		RecTo           string
		RecCc           string
		RecBcc          string
		RecSubject      string
		RecMime         string
		RecFrom         string
		RecCreatedAt    string
		RecUpdatedAt    string
		RecStatus       string
	}{
		RecID:           "rec_id",
		RecEmail:        "rec_email",
		RecTemplateID:   "rec_template_id",
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
)
