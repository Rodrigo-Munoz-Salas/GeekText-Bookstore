-- +goose Up
CREATE TABLE books (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    price FLOAT8 NOT NULL,
    genre TEXT NOT NULL,
    publisher_id UUID NULL REFERENCES publishers(id) ON DELETE SET NULL,
    year_published INT NOT NULL
);

-- +goose Down
DROP TABLE books;
