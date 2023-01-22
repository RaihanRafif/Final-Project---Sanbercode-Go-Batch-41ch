-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE files
(
    id SERIAL PRIMARY KEY,
    filename VARCHAR(255),
    user_id INTEGER,
    class_id INTEGER,
    created_at DATE,
    updated_at DATE
)

-- +migrate StatementEnd