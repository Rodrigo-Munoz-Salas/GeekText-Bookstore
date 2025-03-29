# Bookstore-RESTful-API

This application simulates...

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





