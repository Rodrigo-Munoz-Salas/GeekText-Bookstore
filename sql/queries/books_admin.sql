-- name: CreateBook :one
INSERT INTO books (id, isbn, title, description, price, genre, publisher_id, year_published)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: CreatePublisher :one
INSERT INTO publishers (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetPublisherByName :one
SELECT id FROM publishers WHERE name = $1;

-- name: GetBookByISBN :one
SELECT isbn
FROM books 
WHERE isbn = $1;

