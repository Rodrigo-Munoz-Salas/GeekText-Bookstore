-- +goose Up
CREATE TABLE shopping_cart_books (
    id UUID PRIMARY KEY,
    cart_id UUID REFERENCES shopping_carts(id) ON DELETE CASCADE,
    book_id UUID REFERENCES books(id) ON DELETE CASCADE,
    quantity INT DEFAULT 1
);

-- +goose Down
DROP TABLE shopping_cart_books;