-- +goose Up
CREATE TABLE publishers (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE publishers;