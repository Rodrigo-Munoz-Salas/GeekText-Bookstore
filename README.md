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


---


## Book Rating and Commenting Feature (by Lily Meilan)

This feature allows users to create ratings and comments and also lets them retrieve the average rating and a list of comments for a given book.

### API Endpoints

All endpoints are accessed through the base URL: http://localhost:8080/v1

---

### Create a Rating

**POST** `/rating`
Creates a rating given a Book ID, a rating (must be from 1 - 5, inclusive), and a User ID.

**Example:** http://localhost:8080/v1/rating
{
    "book_id": "08d2132d-6851-42ce-b86d-a6dd97a1e27a",
    "rating": 5,
    "user_id": "abef78ce-85d7-446e-9b90-87ea29252823"
}

---

### Create a Comment

**POST** `/comments`
Creates a comment given a Book ID, a comment, and a User ID.

**Example:**: http://localhost:8080/v1/comments
{
    "book_id": "08d2132d-6851-42ce-b86d-a6dd97a1e27a",
    "comment": "It was amazing!",
    "user_id": "abef78ce-85d7-446e-9b90-87ea29252823"
}

---

### Get the Average Rating of a Book

**GET** `/rating/{bookID}`
Returns the average rating for a given book.

**Example:** http://localhost:8080/v1/rating/08d2132d-6851-42ce-b86d-a6dd97a1e27a

---

### List all Comments for a Book

**GET** `/comments/{bookID}`
Returns a list of all the comments for a given book.

**Example:** http://localhost:8080/v1/comments/08d2132d-6851-42ce-b86d-a6dd97a1e27a


---


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





