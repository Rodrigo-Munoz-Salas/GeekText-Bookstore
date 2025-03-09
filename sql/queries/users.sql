-- name: CreateUser :one
INSERT INTO users (id, username, password_hash, name, email, home_address)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: CreateShoppingCart :one
INSERT INTO shopping_carts (id, user_id)
VALUES ($1, $2)
RETURNING *;
