// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: shopping_cart_books.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addBookToCart = `-- name: AddBookToCart :exec
INSERT INTO shopping_cart_books (id, cart_id, book_id, quantity)
VALUES ($1, $2, $3, $4)
ON CONFLICT (cart_id, book_id) 
DO UPDATE SET quantity = shopping_cart_books.quantity + $4
RETURNING id, cart_id, book_id, quantity
`

type AddBookToCartParams struct {
	ID       uuid.UUID
	CartID   uuid.UUID
	BookID   uuid.UUID
	Quantity sql.NullInt32
}

func (q *Queries) AddBookToCart(ctx context.Context, arg AddBookToCartParams) error {
	_, err := q.db.ExecContext(ctx, addBookToCart,
		arg.ID,
		arg.CartID,
		arg.BookID,
		arg.Quantity,
	)
	return err
}

const getCartSubtotalByUserID = `-- name: GetCartSubtotalByUserID :one
SELECT COALESCE(SUM(b.price * scb.quantity), 0.0)::float8 AS subtotal
FROM shopping_cart_books scb
JOIN books b ON scb.book_id = b.id
JOIN shopping_carts sc ON scb.cart_id = sc.id
WHERE sc.user_id = $1
`

func (q *Queries) GetCartSubtotalByUserID(ctx context.Context, userID uuid.UUID) (float64, error) {
	row := q.db.QueryRowContext(ctx, getCartSubtotalByUserID, userID)
	var subtotal float64
	err := row.Scan(&subtotal)
	return subtotal, err
}

const getShoppingCartByUserID = `-- name: GetShoppingCartByUserID :one
SELECT id
FROM shopping_carts
WHERE user_id = $1
`

func (q *Queries) GetShoppingCartByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getShoppingCartByUserID, userID)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
