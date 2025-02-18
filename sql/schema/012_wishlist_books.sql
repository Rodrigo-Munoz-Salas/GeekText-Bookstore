-- +goose Up
CREATE TABLE wishlist_books (
    id UUID PRIMARY KEY,
    wishlist_id UUID NOT NULL REFERENCES wishlists(id) ON DELETE CASCADE,
    book_id UUID NOT NULL REFERENCES books(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE wishlist_books;