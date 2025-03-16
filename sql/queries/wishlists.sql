-- name: CreateWishlist :one
INSERT INTO wishlists (id, user_id, list_name)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, list_name) DO NOTHING
RETURNING *;

-- name: CountUserWishlists :one
SELECT COUNT(*) FROM wishlists WHERE user_id = $1;

-- name: GetUserIDByWishlistID :one
SELECT user_id FROM wishlists WHERE id = $1;
