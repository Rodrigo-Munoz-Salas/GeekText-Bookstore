-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    name TEXT,
    email TEXT UNIQUE,
    home_address TEXT
);

-- +goose Down
DROP TABLE users;