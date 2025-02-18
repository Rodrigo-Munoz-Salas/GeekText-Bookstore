-- +goose Up
CREATE TABLE authors (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    biography TEXT,
    publisher_id UUID REFERENCES publishers(id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE authors;