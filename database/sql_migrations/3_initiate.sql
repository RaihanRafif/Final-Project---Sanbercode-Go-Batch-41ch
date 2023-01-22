-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE mark
(
    id SERIAL PRIMARY KEY,
    class_id INTEGER,
    mark INTEGER,
    student_id INTEGER,
    created_at DATE,
    updated_at DATE
)

-- +migrate StatementEnd