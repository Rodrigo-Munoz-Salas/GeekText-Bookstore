-- name: AddBookToWishlist :one
INSERT INTO wishlist_books (id, wishlist_id, book_id)
VALUES ($1, $2, $3)
ON CONFLICT (wishlist_id, book_id) DO NOTHING
RETURNING *;

-- name: GetBookToDelete :one
SELECT 1 FROM wishlist_books WHERE wishlist_id = $1 AND book_id = $2 LIMIT 1;

-- name: DeleteBookFromWishlist :exec
DELETE FROM wishlist_books WHERE wishlist_id = $1 AND book_id = $2;

-- name: GetWishlistBooksByWishlistID :many
SELECT b.id, b.isbn, b.title, b.description, b.price, b.genre, b.publisher_id, b.year_published
FROM books b
JOIN wishlist_books wb ON b.id = wb.book_id
WHERE wb.wishlist_id = $1;