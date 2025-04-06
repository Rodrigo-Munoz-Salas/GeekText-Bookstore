# Bookstore-RESTful-API

This is a RESTful API for a ficticious bookstore called GeekText

# Run App

The vendor folder already has all the required modules

## Book Browsing and Sorting Feature (by Pierson Mandell)

This feature allows users to browse books by genre, rating, or top-seller status, and lets administrators apply discounts to books by publisher.

### API Endpoints

All endpoints are accessed through the base URL: http://localhost:8080/v1

---

### Browse Books by Genre

**GET** `/books/genre/{genre}`  
Returns a list of books that match the specified genre.

**Example:** http://localhost:8080/v1/books/genre/Fantasy

---

### Browse Books by Rating

**GET** `/books/rating/{rating}`  
Returns books with an average rating greater than or equal to the specified rating.

**Example:** http://localhost:8080/v1/books/rating/4

---

### Get Top 10 Best-Selling Books

**GET** `/books/top-sellers`  
Returns the 10 most recently published books.

**Example:** http://localhost:8080/v1/books/top-sellers

---

### Apply Discount to Publisher’s Books

**PUT** `/books/discount?discount={discount}&publisher_id={publisher_id}`  
Applies a discount (e.g., 0.20 for 20%) to all books from the given publisher.

**Example:** http://localhost:8080/v1/books/discount?discount=0.20&publisher_id=660e8400-e29b-41d4-a716-446655440001





## Wishlist Management System (by Rodrigo Munoz)

This feature allows users to create their wishlists and populate them with books, as well as moving a book from a wishlist to their shopping cart.

### API Endpoints

All endpoints are accessed through the base URL: http://localhost:8080/v1

---

### Create a Wishlist

Note: All UUIDs are unique. You should replace those with the IDs you generate.

**POST** `/books/wishlists`
JSON Body:
{
    "user_id": "{user_UUID}",
    "list_name": "{wishlist_name}"
}

Creates a wishlist with the provided name for the given user. A Maximum of 3 wishlists can be created for each user.

**Example:** http://localhost:8080/v1/wishlists
{
    "user_id": "7b0e39e0-f1a5-43c8-a2d4-c661e562a3fe",
    "list_name": "My Wishlist"
}

---

### Add a Book to a Wishlist

**POST** `/wishlist_books`
JSON Body:
{
  "wishlist_id": "{wishlist_UUID}",
  "book_id": "{book_UUID}"
}

Adds a the provided book to the given wishlist.


**Example:** http://localhost:8080/v1/wishlist_books
{
  "wishlist_id": "c2acaffc-8603-4f21-8105-0ac4e0392061",
  "book_id": "03d9d733-d089-463b-a53d-9531ca69a758"
}


---

### Remove a Book from a Wishlist

**DELETE** `/wishlist_books/{book_id}`  
JSON Body:
{
  "wishlist_id": "c2acaffc-8603-4f21-8105-0ac4e0392061",
  "to_shopping_cart": "yes" // optional field
}

Removes the given book from the provided wishlist. The user can automatically add it to the shopping cart of the optional field is sent.

**Example:** http://localhost:8080//wishlist_books/03d9d733-d089-463b-a53d-9531ca69a758
{
  "wishlist_id": "c2acaffc-8603-4f21-8105-0ac4e0392061",
  "to_shopping_cart": "yes"
}

---

### List all Books from a Wishlist

**GET** `/wishlist_books`
JSON Body:
{
  "wishlist_id": "{wishlist_UUID}"
}

Returns a list of book objects that belong to the given wishlist

**Example:** http://localhost:8080/v1/wishlist_books
{
  "wishlist_id": "c2acaffc-8603-4f21-8105-0ac4e0392061"
}

---
## Book Details Feature (by Arantza Mendoza)

This feature allows users to create a book, retrieve a book by ISBN, lets and administrator create an author and retrieves books by author.

### API Endpoints

All endpoints are accessed through the base URL: http://localhost:8080/v1

---
### Create Book

**POST** ‘/book_admin/’  
Creates a book object in the bookstore system.

JSON Body: {"isbn": "{isbn}","title": "{title}","description": "{description}","price": "{price}","genre": "{genre}","publisher_name": "{publisher_name}",year_published": "{year_published}","copies_sold": "{copies_sold}"}

Creates a new book object with the provided details in the bookstore system, checking if the publisher exists in the database, and creating the publisher if not.

**Example:** http://localhost:8080/v1/book_admin/{"isbn": "978-3-16-148410-0","title": "The Great Book","description": "A captivating journey of knowledge","price": 19.99,"genre": "Non-Fiction","publisher_name": "Great Publishing","year_published": 2023,"copies_sold": 1000}

---
### Retrieve Book by ISBN

**GET** ‘/book_admin/’  
Retrieves a book object from the bookstore system using its ISBN.

JSON body: {"isbn": "978-3-16-148410-0"}

Retrieves the details of a book from the database using its ISBN. If no book is found, it returns a 404 error.

**Example:** http://localhost:8080/v1/book_admin/{"isbn": "978-3-16-148410-0"}

__ 

### Admin can create an Author

**POST** ‘/book_admin/author’  
Lets an admin create an author object in the system associated with a publisher.

JSON body: {"first_name": "{first_name}","last_name": "{last_name}","biography": "{biography}","publisher_id": "{publisher_id}"}

Creates a new author object in the system, linking the author with a publisher based on the provided publisher ID. If the publisher does not exist, an error is returned.

**Example:** http://localhost:8080/v1/book_admin/author/{"first_name": "John","last_name": "Doe","biography": "John Doe is an acclaimed author of fictional works.","publisher_id": "8c9f7d25-8d0d-4e7d-9e63-847ae15f5e4a"}

__

### Retrieve Books by Author

**GET** ‘/book_admin/author/{author_id}’
Retrieves a list of books written by a specific author.

JSON body: {"author_id": "{author_id}"}

Retrieves a list of books associated with a given author by their unique author ID. If no books are found, it returns a 404 error.

**Example:** http://localhost:8080/v1/book_admin/author/{"author_id": "d600760d-e20e-499c-b7a1-17c448995ada"}

