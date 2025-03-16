-- +goose Up
CREATE TABLE books (
    id UUID PRIMARY KEY,
    isbn VARCHAR(13) UNIQUE,
    title TEXT NOT NULL,
    description TEXT,
    price DECIMAL(6,2) NOT NULL,
    genre TEXT NOT NULL,
    publisher_id UUID REFERENCES publishers(id) ON DELETE SET NULL,
    year_published INT NOT NULL
);

-- +goose Down
DROP TABLE books;