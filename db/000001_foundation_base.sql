-- BEGIN core_module
    CREATE TABLE core_module (
        mod_id                  TEXT NOT NULL,
        mod_name                TEXT NOT NULL,
        mod_description         TEXT NOT NULL,
        mod_order               INTEGER NOT NULL,
        mod_parent_id           TEXT,
        mod_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        mod_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        mod_status              core_dm_record_status
    );

    ALTER TABLE ONLY core_module ADD CONSTRAINT
        core_module_pk PRIMARY KEY (mod_id);

    CALL create_fk('core_module', 'mod_parent_id', 'core_module', 'mod_id', 'padre');
-- END core_module


-- BEGIN core_action
    CREATE TABLE core_action (
        act_module_id           TEXT NOT NULL,
        act_action_id           TEXT NOT NULL,
        act_name                TEXT NOT NULL,
        act_description         TEXT NOT NULL,
        act_order               INTEGER NOT NULL,
        act_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        act_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        act_status              core_dm_record_status
    );

    ALTER TABLE ONLY core_action ADD CONSTRAINT
        core_action_pk PRIMARY KEY (act_module_id, act_action_id);

    CALL create_fk('core_action', 'act_module_id', 'core_module', 'mod_id');
-- END core_action


-- BEGIN core_action_route
    CREATE TABLE core_action_route (
        acr_module_id           TEXT NOT NULL,
        acr_action_id           TEXT NOT NULL,
        acr_method              TEXT NOT NULL,
        acr_route               TEXT NOT NULL
    );

    ALTER TABLE ONLY core_action_route ADD CONSTRAINT
        core_action_route_pk PRIMARY KEY (acr_module_id, acr_action_id, acr_method, acr_route);

    CALL create_fk('core_action_route', 'acr_module_id,acr_action_id', 'core_action', 'act_module_id,act_action_id');
-- END core_action_route


-- BEGIN core_role
    CREATE TABLE core_role (
        rol_id                  TEXT NOT NULL,
        rol_name                TEXT NOT NULL,
        rol_description         TEXT NOT NULL,
        rol_icon                TEXT NOT NULL,
        rol_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        rol_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        rol_status              core_dm_record_status
    );

    ALTER TABLE ONLY core_role ADD CONSTRAINT
        core_role_pk PRIMARY KEY (rol_id);
-- END core_role


-- BEGIN core_privilege
    CREATE TABLE core_privilege(
        pri_role_id             TEXT NOT NULL,
        pri_module_id           TEXT NOT NULL,
        pri_action_id           TEXT NOT NULL
    );

    ALTER TABLE ONLY core_privilege ADD CONSTRAINT
        core_privilege_pk PRIMARY KEY (pri_role_id, pri_module_id, pri_action_id);

    CALL create_fk('core_privilege', 'pri_role_id', 'core_role', 'rol_id');

    CALL create_fk('core_privilege', 'pri_module_id,pri_action_id', 'core_action', 'act_module_id,act_action_id');
-- END core_privilege

-- BEGIN core_user
    CREATE TABLE core_user (
        use_id                  UUID DEFAULT uuid_generate_v1mc() NOT NULL,
        use_name                TEXT NOT NULL,
        use_middle_name         TEXT NOT NULL,
        use_last_name           TEXT NOT NULL,
        use_email               TEXT NOT NULL,
        use_username            TEXT NOT NULL,
        use_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        use_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        use_status              core_dm_record_status
    );

    ALTER TABLE ONLY core_user ADD CONSTRAINT
        core_user_pk PRIMARY KEY (use_id);

    ALTER TABLE ONLY core_user ADD CONSTRAINT
        core_user_uq_use_email UNIQUE (use_email);

    ALTER TABLE ONLY core_user ADD CONSTRAINT
        core_user_uq_use_username UNIQUE (use_username);
-- END core_user


-- BEGIN core_password
    CREATE DOMAIN core_dm_password_type
        AS TEXT NOT NULL
        CHECK (VALUE IN ('Scrypt2017'));

    CREATE TABLE core_password (
        pas_id                  BIGSERIAL NOT NULL,
        pas_data                JSON DEFAULT '{}'::JSON NOT NULL,
        pas_type                core_dm_password_type NOT NULL,
        pas_user_id             UUID NOT NULL,
        pas_invalidation_date   TIMESTAMP WITH TIME ZONE,
        pas_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        pas_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        pas_status              core_dm_record_status
    );

    ALTER TABLE ONLY core_password ADD CONSTRAINT
        core_password_pk PRIMARY KEY (pas_id);

    CALL create_fk('core_password', 'pas_user_id', 'core_user', 'use_id');

    CREATE INDEX core_password_idx_pas_user_id ON core_password USING HASH (pas_user_id);
-- END core_password


-- BEGIN core_user_role
    CREATE TABLE core_user_role(
        usr_user_id             UUID NOT NULL,
        usr_role_id             TEXT NOT NULL
    );

    ALTER TABLE ONLY core_user_role ADD CONSTRAINT
        core_user_role_pk PRIMARY KEY (usr_user_id, usr_role_id);

    CALL create_fk('core_user_role', 'usr_user_id', 'core_user', 'use_id');

    CALL create_fk('core_user_role', 'usr_role_id', 'core_role', 'rol_id');
-- END core_user_role


-- BEGIN core_nav_element
    CREATE TABLE core_nav_element(
        nae_id                  TEXT NOT NULL,
        nae_name                TEXT NOT NULL,
        nae_description         TEXT NOT NULL,
        nae_icon                TEXT NOT NULL,
        nae_order               INTEGER NOT NULL,
        nae_url                 TEXT NOT NULL,
        nae_parent_id           TEXT,
        nae_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        nae_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        nae_status              core_dm_record_status
    );

    ALTER TABLE ONLY core_nav_element ADD CONSTRAINT
        core_nav_element_pk PRIMARY KEY (nae_id);

    CALL create_fk('core_nav_element', 'nae_parent_id', 'core_nav_element', 'nae_id', 'parent', 'ON DELETE SET NULL');
-- END core_nav_element


-- BEGIN core_nav_element_role
    CREATE TABLE core_nav_element_role(
        ner_nav_element_id      TEXT NOT NULL,
        ner_role_id             TEXT NOT NULL
    );

    ALTER TABLE ONLY core_nav_element_role ADD CONSTRAINT
        core_role_nav_element_pk PRIMARY KEY (ner_role_id, ner_nav_element_id);

    CALL create_fk('core_nav_element_role', 'ner_nav_element_id', 'core_nav_element', 'nae_id');

    CALL create_fk('core_nav_element_role', 'ner_role_id', 'core_role', 'rol_id');
-- END core_nav_element_role


-- BEGIN core_account_access
    CREATE DOMAIN core_dm_account_access_type
        AS TEXT NOT NULL
        CHECK (VALUE IN ('RestoreAccount'));

    CREATE TABLE core_account_access (
        aca_id                  TEXT NOT NULL,
        aca_type                core_dm_account_access_type NOT NULL,
        aca_user_id             UUID NOT NULL,
        aca_expiration_date     TIMESTAMP WITH TIME ZONE DEFAULT (now() + '2 days'::interval) NOT NULL,
        aca_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        aca_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        aca_status              core_dm_record_status
    );

    ALTER TABLE ONLY core_account_access ADD CONSTRAINT
        core_access_account_pk PRIMARY KEY (aca_id);

    CALL create_fk('core_account_access', 'aca_user_id', 'core_user', 'use_id');
-- END core_account_access


-- BEGIN core_v_privilege_role
    CREATE VIEW core_v_privilege_role AS
        SELECT
            pri_role_id   prr_role_id,
            pri_module_id prr_module_id,
            pri_action_id prr_action_id,
            acr_method    prr_method,
            acr_route     prr_route
        FROM core_role
        INNER JOIN core_privilege
            ON pri_role_id = rol_id
        INNER JOIN core_action
            ON act_module_id  = pri_module_id
            AND act_action_id = pri_action_id
        LEFT JOIN core_action_route
            ON acr_module_id  = act_module_id
            AND acr_action_id = act_action_id
        WHERE rol_status = 2
        AND act_status = 2;
-- END core_v_privilege_role
