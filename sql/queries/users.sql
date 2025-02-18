-- name: CreateUser :one
INSERT INTO users (id, username, password_hash, name, email, home_address)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;