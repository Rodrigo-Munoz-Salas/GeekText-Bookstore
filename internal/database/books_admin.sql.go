// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: books_admin.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (id, first_name, last_name, biography, publisher_id)
VALUES (gen_random_uuid(), $1, $2, $3, $4)
RETURNING id
`

type CreateAuthorParams struct {
	FirstName   string
	LastName    string
	Biography   sql.NullString
	PublisherID uuid.NullUUID
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createAuthor,
		arg.FirstName,
		arg.LastName,
		arg.Biography,
		arg.PublisherID,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createBook = `-- name: CreateBook :one
INSERT INTO books (id, isbn, title, description, price, genre, publisher_id, year_published)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, isbn, title, description, price, genre, publisher_id, year_published
`

type CreateBookParams struct {
	ID            uuid.UUID
	Isbn          string
	Title         string
	Description   sql.NullString
	Price         string
	Genre         string
	PublisherID   uuid.NullUUID
	YearPublished int32
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.ID,
		arg.Isbn,
		arg.Title,
		arg.Description,
		arg.Price,
		arg.Genre,
		arg.PublisherID,
		arg.YearPublished,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Isbn,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.Genre,
		&i.PublisherID,
		&i.YearPublished,
	)
	return i, err
}

const createPublisher = `-- name: CreatePublisher :one
INSERT INTO publishers (id, name)
VALUES ($1, $2)
RETURNING id, name
`

type CreatePublisherParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) CreatePublisher(ctx context.Context, arg CreatePublisherParams) (Publisher, error) {
	row := q.db.QueryRowContext(ctx, createPublisher, arg.ID, arg.Name)
	var i Publisher
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getBookByISBN = `-- name: GetBookByISBN :one
SELECT id, isbn, title, description, price, genre, publisher_id, year_published 
FROM books 
WHERE isbn = $1
`

func (q *Queries) GetBookByISBN(ctx context.Context, isbn string) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBookByISBN, isbn)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Isbn,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.Genre,
		&i.PublisherID,
		&i.YearPublished,
	)
	return i, err
}

const getPublisherByName = `-- name: GetPublisherByName :one
SELECT id FROM publishers WHERE name = $1
`

func (q *Queries) GetPublisherByName(ctx context.Context, name string) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getPublisherByName, name)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
