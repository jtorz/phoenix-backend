
--drop database phoenix;
--create database phoenix;
CREATE EXTENSION IF NOT EXISTS intarray WITH SCHEMA public;

COMMENT ON EXTENSION intarray IS 'functions, operators, and index support for 1-D arrays of integers';

CREATE EXTENSION IF NOT EXISTS tablefunc WITH SCHEMA public;

COMMENT ON EXTENSION tablefunc IS 'functions that manipulate whole tables, including crosstab';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';

CREATE DOMAIN fnd_dm_record_status
    AS SMALLINT DEFAULT 1 NOT NULL
    CHECK (VALUE IN (0, 1, 2, 3));

-- BEGIN fnd_module
    CREATE TABLE fnd_module (
        mod_id                  TEXT NOT NULL,
        mod_name                TEXT NOT NULL,
        mod_description         TEXT NOT NULL,
        mod_order               INTEGER NOT NULL,
        mod_parent_id           TEXT,
        mod_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        mod_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        mod_status              fnd_dm_record_status
    );

    ALTER TABLE ONLY fnd_module ADD CONSTRAINT
        fnd_module_pk PRIMARY KEY (mod_id);

    ALTER TABLE ONLY fnd_module ADD CONSTRAINT
        fnd_module_fk_fnd_module FOREIGN KEY (mod_parent_id)
        REFERENCES fnd_module(mod_id);
-- END fnd_module


-- BEGIN fnd_action
    CREATE TABLE fnd_action (
        act_module_id           TEXT NOT NULL,
        act_action_id           TEXT NOT NULL,
        act_name                TEXT NOT NULL,
        act_description         TEXT NOT NULL,
        act_order               INTEGER NOT NULL,
        act_route               TEXT NOT NULL,
        act_method              TEXT NOT NULL,
        act_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        act_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        act_status              fnd_dm_record_status
    );

    ALTER TABLE ONLY fnd_action ADD CONSTRAINT
        fndtaction_pk PRIMARY KEY (act_module_id, act_action_id);

    ALTER TABLE ONLY fnd_action ADD CONSTRAINT
        fndtaction_fk_fnd_module FOREIGN KEY (act_module_id)
        REFERENCES fnd_module(mod_id);
-- END fnd_action


-- BEGIN fnd_role
    CREATE TABLE fnd_role (
        rol_id                  TEXT NOT NULL,
        rol_name                TEXT NOT NULL,
        rol_description         TEXT NOT NULL,
        rol_icon                TEXT NOT NULL,
        rol_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        rol_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        rol_status              fnd_dm_record_status
    );

    ALTER TABLE ONLY fnd_role ADD CONSTRAINT
        fndtrole_pk PRIMARY KEY (rol_id);
-- END fnd_role


-- BEGIN fnd_privilege
    CREATE TABLE fnd_privilege(
        pri_role_id             TEXT NOT NULL,
        pri_module_id           TEXT NOT NULL,
        pri_action_id           TEXT NOT NULL
    );

    ALTER TABLE ONLY fnd_privilege ADD CONSTRAINT
        fndtprivilege_pk PRIMARY KEY (pri_role_id, pri_module_id, pri_action_id);

    ALTER TABLE ONLY fnd_privilege ADD CONSTRAINT
        fndtprivilege_fk_fndtrole FOREIGN KEY (pri_role_id)
        REFERENCES fnd_role(rol_id);

    ALTER TABLE ONLY fnd_privilege ADD CONSTRAINT
        fndtprivilege_fk_fndtaction FOREIGN KEY (pri_module_id, pri_action_id)
        REFERENCES fnd_action(act_module_id, act_action_id);
-- END fnd_privilege

-- BEGIN fnd_user
    CREATE TABLE fnd_user (
        use_id                  UUID DEFAULT uuid_generate_v1mc() NOT NULL,
        use_name                TEXT NOT NULL,
        use_middle_name         TEXT NOT NULL,
        use_last_name           TEXT NOT NULL,
        use_email               TEXT NOT NULL,
        use_username            TEXT NOT NULL,
        use_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        use_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        use_status              fnd_dm_record_status
    );

    ALTER TABLE ONLY fnd_user ADD CONSTRAINT
        fndtuser_pk PRIMARY KEY (use_id);

    ALTER TABLE ONLY fnd_user ADD CONSTRAINT
        fndtuser_uq_use_email UNIQUE (use_email);

    ALTER TABLE ONLY fnd_user ADD CONSTRAINT
        fndtuser_uq_use_username UNIQUE (use_username);
-- END fnd_user


-- BEGIN fnd_password
    CREATE DOMAIN fnd_dm_password_type
        AS TEXT NOT NULL
        CHECK (VALUE IN ('Scrypt2017'));

    CREATE TABLE fnd_password (
        pas_id                  BIGSERIAL NOT NULL,
        pas_data                JSON DEFAULT '{}'::JSON NOT NULL,
        pas_type                fnd_dm_password_type NOT NULL,
        pas_user_id             UUID NOT NULL,
        pas_invalidation_date   TIMESTAMP WITH TIME ZONE,
        pas_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        pas_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        pas_status              fnd_dm_record_status
    );

    ALTER TABLE ONLY fnd_password ADD CONSTRAINT
        fndtpassword_pk PRIMARY KEY (pas_id);

    ALTER TABLE ONLY fnd_password ADD CONSTRAINT
        fndtpassword_fk_fndtuser FOREIGN KEY (pas_user_id)
        REFERENCES fnd_user(use_id);

    CREATE INDEX fndtpassword_idx_pas_user_id ON fnd_password USING HASH (pas_user_id);
-- END fnd_password


-- BEGIN fnd_user_role
    CREATE TABLE fnd_user_role(
        usr_user_id             UUID NOT NULL,
        usr_role_id             TEXT NOT NULL
    );

    ALTER TABLE ONLY fnd_user_role ADD CONSTRAINT
        fndtuser_role_pk PRIMARY KEY (usr_user_id, usr_role_id);

    ALTER TABLE ONLY fnd_user_role ADD CONSTRAINT
        fndtuser_role_fk_fndtuser FOREIGN KEY (usr_user_id)
        REFERENCES fnd_user(use_id);

    ALTER TABLE ONLY fnd_user_role ADD CONSTRAINT
        fndtuser_role_fk_fndtrole FOREIGN KEY (usr_role_id)
        REFERENCES fnd_role(rol_id);
-- END fnd_user_role


-- BEGIN fnd_navigator
    CREATE TABLE fnd_navigator(
        nav_id                  TEXT NOT NULL,
        nav_name                TEXT NOT NULL,
        nav_description         TEXT NOT NULL,
        nav_icon                TEXT NOT NULL,
        nav_order               TEXT NOT NULL,
        nav_url                 TEXT NOT NULL,
        nav_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        nav_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        nav_status              fnd_dm_record_status
    );

    ALTER TABLE ONLY fnd_navigator ADD CONSTRAINT
        fndtnavigator_pk PRIMARY KEY (nav_id);
-- END fnd_navigator


-- BEGIN fnd_role_navigator
    CREATE TABLE fnd_role_navigator(
        ron_role_id             TEXT NOT NULL,
        ron_navigator_id        TEXT NOT NULL
    );

    ALTER TABLE ONLY fnd_role_navigator ADD CONSTRAINT
        fndtrole_navigator_pk PRIMARY KEY (ron_role_id, ron_navigator_id);

    ALTER TABLE ONLY fnd_role_navigator ADD CONSTRAINT
        fndtrole_navigator_fk_fndtrole FOREIGN KEY (ron_role_id)
        REFERENCES fnd_role(rol_id);

    ALTER TABLE ONLY fnd_role_navigator ADD CONSTRAINT
        fndtrole_navigator_fk_fndtnavigator FOREIGN KEY (ron_navigator_id)
        REFERENCES fnd_navigator(nav_id);
-- END fnd_role_navigator


-- BEGIN fnd_account_access
    CREATE DOMAIN fnd_dm_account_access_type
        AS TEXT NOT NULL
        CHECK (VALUE IN ('RestoreAccount'));

    CREATE TABLE fnd_account_access (
        aca_id                  TEXT NOT NULL,
        aca_type                fnd_dm_account_access_type NOT NULL,
        aca_user_id             UUID NOT NULL,
        aca_expiration_date     TIMESTAMP WITH TIME ZONE DEFAULT (now() + '2 days'::interval) NOT NULL,
        aca_created_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        aca_updated_at          TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
        aca_status              fnd_dm_record_status
    );

    ALTER TABLE ONLY fnd_account_access ADD CONSTRAINT
        fndtaccess_account_pk PRIMARY KEY (aca_id);

    ALTER TABLE ONLY fnd_account_access ADD CONSTRAINT
        fndtaccess_account_fk_fndtuser FOREIGN KEY (aca_user_id)
        REFERENCES fnd_user(use_id);
-- END fnd_account_access


-- BEGIN fnd_v_privilege_role
    CREATE VIEW fnd_v_privilege_role AS
        SELECT
            pri_role_id   prr_role_id,
            pri_module_id prr_module_id,
            pri_action_id prr_action_id,
            act_route     prr_route,
            act_method    prr_method
        FROM fnd_role
        INNER JOIN fnd_privilege
            ON pri_role_id = rol_id
        INNER JOIN fnd_action
            ON act_module_id  = pri_module_id
            AND act_action_id = pri_action_id
        WHERE rol_status = 2
        AND act_status = 2;
-- END fnd_v_privilege_role