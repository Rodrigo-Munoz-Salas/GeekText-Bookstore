-- name: GetShoppingCartByUserID :one
SELECT id
FROM shopping_carts
WHERE user_id = $1;

-- name: AddBookToCart :exec
INSERT INTO shopping_cart_books (id, cart_id, book_id, quantity)
VALUES ($1, $2, $3, $4)
ON CONFLICT (cart_id, book_id) 
DO UPDATE SET quantity = shopping_cart_books.quantity + $4
RETURNING *;
