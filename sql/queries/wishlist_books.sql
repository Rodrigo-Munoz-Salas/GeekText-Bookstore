-- name: AddBookToWishlist :one
INSERT INTO wishlist_books (id, wishlist_id, book_id)
VALUES ($1, $2, $3)
ON CONFLICT (wishlist_id, book_id) DO NOTHING
RETURNING *;