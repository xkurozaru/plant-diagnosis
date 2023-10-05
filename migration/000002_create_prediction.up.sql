BEGIN;

CREATE TABLE IF NOT EXISTS prediction_model_entities (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    name TEXT NOT NULL,
    network_name TEXT NOT NULL,
    param_path TEXT NOT NULL,
);

CREATE TABLE IF NOT EXISTS prediction_label_entities (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    prediction_model_id VARCHAR(26) NOT NULL REFERENCES prediction_model_entities(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    index UNSIGNED INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS prediction_result_entities (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_id VARCHAR(26) NOT NULL REFERENCES user_entities(id) ON DELETE CASCADE,
    prediction_model_id VARCHAR(26) NOT NULL REFERENCES prediction_model_entities(id) ON DELETE CASCADE,
    result TEXT NOT NULL,
    file_path TEXT
);

COMMIT;
