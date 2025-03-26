-- name: GetBooksByGenre :many
SELECT id, title, description, price, genre, publisher_id, year_published 
FROM books 
WHERE genre = $1;

-- name: GetTopSellingBooks :many
SELECT b.id, b.title, b.description, b.price, b.genre, b.publisher_id, b.year_published 
FROM books b
JOIN book_authors ba ON b.id = ba.book_id
JOIN authors a ON ba.author_id = a.id
ORDER BY b.year_published DESC
LIMIT 10;

-- name: GetBooksByRating :many
SELECT b.id, b.title, b.description, b.price, b.genre, b.publisher_id, b.year_published, AVG(r.rating) AS average_rating
FROM books b
JOIN ratings r ON b.id = r.book_id
GROUP BY b.id
HAVING AVG(r.rating) >= $1;

-- name: ApplyDiscountToPublisher :exec
UPDATE books 
SET price = price * (1 - sqlc.arg(discount_percent)::FLOAT8) 
WHERE publisher_id = sqlc.arg(publisher_id)::UUID;
