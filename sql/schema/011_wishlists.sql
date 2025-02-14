-- +goose Up
CREATE TABLE wishlists (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    list_name TEXT NOT NULL CHECK (TRIM(list_name) <> ''),
    UNIQUE (user_id, list_name)
);

-- +goose Down
DROP TABLE wishlists;