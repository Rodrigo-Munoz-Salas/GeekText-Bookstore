-- +goose Up
ALTER TABLE books ADD COLUMN copies_sold INT DEFAULT 10 NOT NULL;

-- +goose Down
ALTER TABLE books DROP COLUMN copies_sold;
