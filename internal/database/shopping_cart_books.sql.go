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

const checkBookInCart = `-- name: CheckBookInCart :one
SELECT EXISTS(
    SELECT 1
    FROM shopping_cart_books
    WHERE cart_id = $1 AND book_id = $2
) AS exists
`

type CheckBookInCartParams struct {
	CartID uuid.UUID
	BookID uuid.UUID
}

func (q *Queries) CheckBookInCart(ctx context.Context, arg CheckBookInCartParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkBookInCart, arg.CartID, arg.BookID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const deleteBookFromCart = `-- name: DeleteBookFromCart :exec
DELETE FROM shopping_cart_books
WHERE cart_id = $1 AND book_id = $2
`

type DeleteBookFromCartParams struct {
	CartID uuid.UUID
	BookID uuid.UUID
}

func (q *Queries) DeleteBookFromCart(ctx context.Context, arg DeleteBookFromCartParams) error {
	_, err := q.db.ExecContext(ctx, deleteBookFromCart, arg.CartID, arg.BookID)
	return err
}

const getCartBooksByUserID = `-- name: GetCartBooksByUserID :many
SELECT scb.book_id, b.title, b.isbn, COALESCE(b.description, ''), b.price, 
        b.genre, b.publisher_id, b.year_published,
        COALESCE(scb.quantity,0) AS quantity
FROM shopping_cart_books scb   
JOIN books b ON scb.book_id = b.id
JOIN shopping_carts sc ON scb.cart_id = sc.id
WHERE sc.user_id = $1
`

type GetCartBooksByUserIDRow struct {
	BookID        uuid.UUID
	Title         string
	Isbn          string
	Description   string
	Price         string
	Genre         string
	PublisherID   uuid.NullUUID
	YearPublished int32
	Quantity      int32
}

func (q *Queries) GetCartBooksByUserID(ctx context.Context, userID uuid.UUID) ([]GetCartBooksByUserIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getCartBooksByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCartBooksByUserIDRow
	for rows.Next() {
		var i GetCartBooksByUserIDRow
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Isbn,
			&i.Description,
			&i.Price,
			&i.Genre,
			&i.PublisherID,
			&i.YearPublished,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
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
