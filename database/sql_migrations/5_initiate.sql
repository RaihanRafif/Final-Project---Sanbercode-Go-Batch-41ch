-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE teacher
(
    id SERIAL PRIMARY KEY,
    phone BIGINT UNIQUE,
    password VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    username VARCHAR(255) UNIQUE,
    role VARCHAR(255),
    created_at DATE,
    updated_at DATE
)

-- +migrate StatementEnd