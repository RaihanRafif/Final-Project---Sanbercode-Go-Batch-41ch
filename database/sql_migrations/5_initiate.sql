-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE member
(
    id SERIAL PRIMARY KEY,
    class_id INTEGER,
    user_id INTEGER,
    created_at DATE,
    updated_at DATE
)

-- +migrate StatementEnd