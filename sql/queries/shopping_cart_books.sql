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

-- name: GetCartSubtotalByUserID :one
SELECT COALESCE(SUM(b.price * scb.quantity), 0.0)::float8 AS subtotal
FROM shopping_cart_books scb
JOIN books b ON scb.book_id = b.id
JOIN shopping_carts sc ON scb.cart_id = sc.id
WHERE sc.user_id = $1;

-- name: GetCartBooksByUserID :many
SELECT scb.book_id, b.title, b.price, COALESCE(scb.quantity,0) AS quantity
FROM shopping_cart_books scb   
JOIN books b ON scb.book_id = b.id
JOIN shopping_carts sc ON scb.cart_id = sc.id
WHERE sc.user_id = $1;
