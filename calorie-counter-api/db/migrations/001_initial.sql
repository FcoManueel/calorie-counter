--# Create a user named `gouser` and run the following commands in your terminal:
--# DB_NAME=calories_dev
--# dropdb $DB_NAME
--# createdb -O gouser $DB_NAME
--# echo 'CREATE EXTENSION "uuid-ossp";' | psql $DB_NAME

-- Users table
CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    role text DEFAULT 'user'::text NOT NULL,
    name text,
    email text NOT NULL,
    password text,
    goal_calories integer
);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_key UNIQUE (email);

-- Intake table
CREATE TABLE intakes (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    name text NOT NULL,
    calories integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    consumed_at timestamp with time zone NOT NULL
);

ALTER TABLE ONLY intakes
    ADD CONSTRAINT intakes_pkey PRIMARY KEY (id);


-- Auth tokens
CREATE TABLE auth_tokens (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    token text
);