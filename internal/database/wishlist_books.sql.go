// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: wishlist_books.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const addBookToWishlist = `-- name: AddBookToWishlist :one
INSERT INTO wishlist_books (id, wishlist_id, book_id)
VALUES ($1, $2, $3)
ON CONFLICT (wishlist_id, book_id) DO NOTHING
RETURNING id, wishlist_id, book_id
`

type AddBookToWishlistParams struct {
	ID         uuid.UUID
	WishlistID uuid.UUID
	BookID     uuid.UUID
}

func (q *Queries) AddBookToWishlist(ctx context.Context, arg AddBookToWishlistParams) (WishlistBook, error) {
	row := q.db.QueryRowContext(ctx, addBookToWishlist, arg.ID, arg.WishlistID, arg.BookID)
	var i WishlistBook
	err := row.Scan(&i.ID, &i.WishlistID, &i.BookID)
	return i, err
}

const deleteBookFromWishlist = `-- name: DeleteBookFromWishlist :exec
DELETE FROM wishlist_books WHERE wishlist_id = $1 AND book_id = $2
`

type DeleteBookFromWishlistParams struct {
	WishlistID uuid.UUID
	BookID     uuid.UUID
}

func (q *Queries) DeleteBookFromWishlist(ctx context.Context, arg DeleteBookFromWishlistParams) error {
	_, err := q.db.ExecContext(ctx, deleteBookFromWishlist, arg.WishlistID, arg.BookID)
	return err
}
