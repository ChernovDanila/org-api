-- +goose Up
CREATE TABLE department (
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(200),
    parentID INTEGER,
    createdAt TIMESTAMPTZ
);

-- +goose Down
SELECT 'down SQL query';
