-- name: CreateBook :one
INSERT INTO books (id, isbn, title, description, price, genre, publisher_id, year_published, copies_sold, author)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: CreatePublisher :one
INSERT INTO publishers (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetPublisherByName :one
SELECT id FROM publishers WHERE name = $1;

-- name: GetAuthorByName :one
SELECT id FROM authors WHERE first_name = $1 AND last_name =$2;

-- name: GetBookByISBN :one
SELECT id, isbn, title, description, price, genre, publisher_id, year_published, copies_sold, author
FROM books 
WHERE isbn = $1;

-- name: CreateAuthor :one
INSERT INTO authors (id, first_name, last_name, biography, publisher_id)
VALUES (gen_random_uuid(), $1, $2, $3, $4)
RETURNING id;

-- name: GetBookIdsByAuthorId :many
SELECT book_id
FROM book_authors
WHERE author_id = $1;

-- name: GetBookDetailsByBookId :one
SELECT id, isbn, title, description, price, genre, publisher_id, year_published, copies_sold, author
FROM books
WHERE id = $1;

