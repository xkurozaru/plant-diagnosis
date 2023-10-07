BEGIN;

CREATE TABLE IF NOT EXISTS user_entities (
    id varchar(26) PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    login_id varchar(32) NOT NULL UNIQUE,
    name TEXT NOT NULL,
    role TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS authentication_entities (
    id varchar(26) PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id varchar(26) NOT NULL,
    hashed_password varchar(60) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_entities(id) ON DELETE CASCADE
);

COMMIT;
