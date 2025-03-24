-- name: CreateUser :one
INSERT INTO users (id, username, password_hash, name, email, home_address)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: CreateShoppingCart :one
INSERT INTO shopping_carts (id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users where username = $1;

-- name: UpdateUserByUsername :one
UPDATE users SET
    password_hash = CASE WHEN $2 != '' THEN $2 ELSE password_hash END,
    name = CASE WHEN $3 != '' THEN $3 ELSE name END,
    home_address = CASE WHEN $4 != '' THEN $4 ELSE home_address END
WHERE username = $1
RETURNING *;

-- name: CreateUserCreditCard :one
INSERT INTO credit_cards (id, user_id, card_number, expiration_date, cvv)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
