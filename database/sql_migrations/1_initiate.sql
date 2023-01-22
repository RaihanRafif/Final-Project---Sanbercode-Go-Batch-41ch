-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE class
(
    id SERIAL PRIMARY KEY,
    topic VARCHAR(255) UNIQUE,
    max_marks INTEGER,
    teacher_id INTEGER,
    description VARCHAR(255),
    created_at DATE,
    updated_at DATE
)

-- +migrate StatementEnd