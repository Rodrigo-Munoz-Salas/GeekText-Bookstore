-- Insert users
INSERT INTO users (id, username, password_hash, name, email, home_address) VALUES
(gen_random_uuid(), 'roary22', 'paws_up', 'Roary', 'roary@fiu.edu', 'FIU'),
(gen_random_uuid(), 'rodms', 'test123', 'Rodrigo', 'rmuno@fiu.edu', 'Florida');

-- Insert publishers
INSERT INTO publishers (id, name) VALUES
(gen_random_uuid(), 'Penguin Books'),
(gen_random_uuid(), 'HarperCollins'),
(gen_random_uuid(), 'Random House');

-- Insert authors
INSERT INTO authors (id, first_name, last_name, biography, publisher_id) VALUES
(gen_random_uuid(), 'George', 'Orwell', 'English novelist and journalist', (SELECT id FROM publishers WHERE name = 'Penguin Books')),
(gen_random_uuid(), 'J.K.', 'Rowling', 'British author, best known for Harry Potter', (SELECT id FROM publishers WHERE name = 'Bloomsbury')),
(gen_random_uuid(), 'Agatha', 'Christie', 'British writer known for mystery novels', (SELECT id FROM publishers WHERE name = 'HarperCollins'));

-- Insert books
INSERT INTO books (id, title, isbn, description, price, genre, publisher_id, year_published, copies_sold) VALUES
(gen_random_uuid(), '1984', '9780451524935','Dystopian social science fiction novel', 14.99, 'Fiction', (SELECT id FROM publishers WHERE name = 'Penguin Books'), 1949, 100),
(gen_random_uuid(), 'Harry Potter and the Sorcerer''s Stone', '9780590353427','Fantasy novel about a young wizard', 19.99, 'Fantasy', (SELECT id FROM publishers WHERE name = 'Bloomsbury'), 1997, 300),
(gen_random_uuid(), 'Murder on the Orient Express', '9780062693662','Detective novel featuring Hercule Poirot', 12.99, 'Mystery', (SELECT id FROM publishers WHERE name = 'HarperCollins'), 1934, 200);

-- Insert book_authors
INSERT INTO book_authors (book_id, author_id) VALUES
((SELECT id FROM books WHERE title = '1984'), (SELECT id FROM authors WHERE first_name = 'George' AND last_name = 'Orwell')),
((SELECT id FROM books WHERE title = 'Harry Potter and the Sorcerer''s Stone'), (SELECT id FROM authors WHERE first_name = 'J.K.' AND last_name = 'Rowling')),
((SELECT id FROM books WHERE title = 'Murder on the Orient Express'), (SELECT id FROM authors WHERE first_name = 'Agatha' AND last_name = 'Christie'));

-- Insert credit cards
INSERT INTO credit_cards (id, user_id, card_number, expiration_date, cvv) VALUES
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'roary22'), '1234567812345678', '2026-12-31', '123'),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'rodms'), '8765432187654321', '2025-11-30', '456');

-- Insert shopping carts
INSERT INTO shopping_carts (id, user_id) VALUES
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'roary22')),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'rodms'));

-- Insert shopping_cart_books
INSERT INTO shopping_cart_books (id, cart_id, book_id, quantity) VALUES
(gen_random_uuid(), (SELECT id FROM shopping_carts WHERE user_id = (SELECT id FROM users WHERE username = 'roary22')), (SELECT id FROM books WHERE title = '1984'), 1),
(gen_random_uuid(), (SELECT id FROM shopping_carts WHERE user_id = (SELECT id FROM users WHERE username = 'rodms')), (SELECT id FROM books WHERE title = 'Harry Potter and the Sorcerer''s Stone'), 2);

-- Insert ratings
INSERT INTO ratings (id, user_id, book_id, rating, created_at) VALUES
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'roary22'), (SELECT id FROM books WHERE title = '1984'), 5, NOW()),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'rodms'), (SELECT id FROM books WHERE title = 'Harry Potter and the Sorcerer''s Stone'), 4, NOW());

-- Insert comments
INSERT INTO comments (id, user_id, book_id, comment, created_at) VALUES
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'roary22'), (SELECT id FROM books WHERE title = '1984'), 'A masterpiece of dystopian fiction!', NOW()),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'rodms'), (SELECT id FROM books WHERE title = 'Harry Potter and the Sorcerer''s Stone'), 'Loved the magic and adventure!', NOW());

-- Insert wishlists
INSERT INTO wishlists (id, user_id, list_name) VALUES
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'roary22'), 'Favorite Books'),
(gen_random_uuid(), (SELECT id FROM users WHERE username = 'rodms'), 'Must Read');

-- Insert wishlist_books
INSERT INTO wishlist_books (id, wishlist_id, book_id) VALUES
(gen_random_uuid(), (SELECT id FROM wishlists WHERE user_id = (SELECT id FROM users WHERE username = 'roary22') AND list_name = 'Favorite Books'), (SELECT id FROM books WHERE title = 'Murder on the Orient Express')),
(gen_random_uuid(), (SELECT id FROM wishlists WHERE user_id = (SELECT id FROM users WHERE username = 'rodms') AND list_name = 'Must Read'), (SELECT id FROM books WHERE title = '1984'));
