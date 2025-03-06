-- name: CreateRating :one
INSERT INTO ratings (id, book_id, rating, user_id, created_at)
VALUES ($1, $2, $3, $4, NOW())
RETURNING *;

-- name: CreateComment :one
INSERT INTO comments (id, book_id, comment, user_id, created_at)
VALUES ($1, $2, $3, $4, NOW())
RETURNING *;

-- name: GetCommentsByBook :many
SELECT id, book_id, comment, user_id, created_at
FROM comments
WHERE book_id = $1
ORDER BY created_at DESC;

-- name: GetAveRatingByBook :one
SELECT COALESCE(AVG(rating), 0) AS average_rating
FROM ratings 
WHERE book_id = $1;