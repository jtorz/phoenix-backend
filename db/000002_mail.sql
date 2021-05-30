
-- BEGIN mail_sender
    CREATE TABLE mail_sender(
        sen_id                  TEXT NOT NULL,
        sen_name                TEXT NOT NULL,
        sen_description         TEXT NOT NULL,
        sen_host                TEXT NOT NULL,
        sen_port                INTEGER NOT NULL,
        sen_user                TEXT NOT NULL,
        sen_password            TEXT NOT NULL,
        sen_from                TEXT NOT NULL,
        sen_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        sen_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        sen_status              fnd_dm_record_status
    );

    ALTER TABLE ONLY mail_sender ADD CONSTRAINT
        mail_sender_pk PRIMARY KEY (sen_id);
-- END mail_sender


-- BEGIN mail_template_type
    CREATE TABLE mail_template_type (
        tet_id         TEXT NOT NULL,
        tet_name       TEXT NOT NULL,
        tet_tags       JSON DEFAULT '[]'::JSON NOT NULL,
        tet_created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        tet_updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        tet_status     fnd_dm_record_status
    );

    ALTER TABLE ONLY mail_template_type ADD CONSTRAINT
        mail_template_type_pk PRIMARY KEY (tet_id);
-- END mail_template_type

-- BEGIN mail_header
    CREATE TABLE mail_header (
        hea_id         SERIAL NOT NULL,
        hea_name       TEXT NOT NULL,
        hea_header     TEXT NOT NULL,
        hea_created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        hea_updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        hea_status     fnd_dm_record_status
    );

    ALTER TABLE ONLY mail_header ADD CONSTRAINT
        mail_header_pk PRIMARY KEY (hea_id);
-- END mail_header

-- BEGIN mail_footer
    CREATE TABLE mail_footer (
        foo_id         SERIAL NOT NULL,
        foo_name       TEXT NOT NULL,
        foo_footer     TEXT NOT NULL,
        foo_created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        foo_updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        foo_status     fnd_dm_record_status
    );

    ALTER TABLE ONLY mail_footer ADD CONSTRAINT
        mail_footer_pk PRIMARY KEY (foo_id);
-- END mail_footer

-- BEGIN mail_template
    CREATE TABLE mail_template (
        tem_type_id     TEXT NOT NULL,
        tem_template_id SERIAL NOT NULL,
        tem_from        TEXT NOT NULL,
        tem_subject     TEXT NOT NULL,
        tem_template    TEXT NOT NULL,
        tem_header_id   INTEGER NOT NULL,
        tem_footer_id   INTEGER NOT NULL,
        tem_sender_id   TEXT NOT NULL,
        tem_created_at  TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        tem_updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        tem_status      fnd_dm_record_status
    );

    ALTER TABLE ONLY mail_template ADD CONSTRAINT
        mail_template_pk PRIMARY KEY (tem_type_id, tem_template_id);

    CALL create_fk('mail_template', 'tem_type_id', 'mail_template_type', 'tet_id');

    CALL create_fk('mail_template', 'tem_header_id', 'mail_header', 'hea_id');

    CALL create_fk('mail_template', 'tem_footer_id', 'mail_footer', 'foo_id');

    CALL create_fk('mail_template', 'tem_sender_id', 'mail_sender', 'sen_id');
-- END mail_template

-- BEGIN mail_b_record
    CREATE TABLE mail_b_record (
        rec_id               BIGSERIAL NOT NULL,
        rec_type_id          TEXT,
        rec_email            TEXT NOT NULL,
        rec_sender_user_id   UUID,
        rec_error            TEXT,
        rec_to               TEXT NOT NULL,
        rec_cc               TEXT,
        rec_bcc              TEXT,
        rec_subject          TEXT NOT NULL,
        rec_mime             TEXT,
        rec_from             TEXT NOT NULL,
        rec_created_at       TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        rec_updated_at       TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        rec_status           fnd_dm_record_status
    );

    ALTER TABLE ONLY mail_b_record ADD CONSTRAINT
        mail_b_record_pk PRIMARY KEY (rec_id);

    CALL create_fk('mail_b_record', 'rec_type_id', 'mail_template_type', 'tet_id');

    CALL create_fk('mail_b_record', 'rec_sender_user_id', 'fnd_user', 'use_id');
-- END mail_b_record