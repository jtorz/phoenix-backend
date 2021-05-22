
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


-- BEGIN create_fk
    CREATE OR REPLACE PROCEDURE create_fk(
        table_           REGCLASS,
        columns_         TEXT,
        foreign_table_   REGCLASS,
        foreign_columns_ TEXT,
        use_             TEXT
    ) LANGUAGE plpgsql
    AS $$
    DECLARE
        cols_ TEXT;
        fcols_ TEXT;
        safe_use_ TEXT;

        clean_columns CURSOR (unsafe_ TEXT) FOR
            SELECT STRING_AGG(col_name, ',')
            FROM (
                SELECT COALESCE(columns.column_name::text, '_UNKNOWN_COLUMN_') col_name
                FROM unnest(string_to_array(unsafe_, ',')) WITH ORDINALITY t(column_name, ord)
                LEFT JOIN information_schema.columns ON trim(t.column_name) = columns.column_name
                ORDER  BY t.ord
            ) v;
    BEGIN
        safe_use_ := regexp_replace(use_, '[^a-zA-Z]', '', 'g');
        IF(safe_use_ IS NOT NULL OR safe_use_ != '') THEN
            safe_use_ := '_' || safe_use_;
        END IF;

	    OPEN  clean_columns(columns_);
        FETCH clean_columns INTO cols_;
        CLOSE clean_columns;

        OPEN  clean_columns(foreign_columns_);
        FETCH clean_columns INTO fcols_;
        CLOSE clean_columns;

        EXECUTE 'ALTER TABLE ONLY ' ||table_|| ' ADD CONSTRAINT ' ||
        table_ || '_fk_' || foreign_table_ || safe_use_ || '  FOREIGN KEY (' || cols_ || '  )' ||
        'REFERENCES ' || foreign_table_ || '(' || fcols_ || ');';
    END; $$
-- END create_fk
