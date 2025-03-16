-- name: AddBookToWishlist :one
INSERT INTO wishlist_books (id, wishlist_id, book_id)
VALUES ($1, $2, $3)
ON CONFLICT (wishlist_id, book_id) DO NOTHING
RETURNING *;

-- name: DeleteBookFromWishlist :exec
DELETE FROM wishlist_books WHERE wishlist_id = $1 AND book_id = $2;