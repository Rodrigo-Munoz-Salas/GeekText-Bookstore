-- +goose Up
ALTER TABLE books ADD COLUMN author TEXT NOT NULL DEFAULT 'Unknown';

-- +goose Down
ALTER TABLE books DROP COLUMN author;
