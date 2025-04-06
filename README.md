# Bookstore-RESTful-API

This is a RESTful API for a ficticious bookstore called GeekText

# Run App

The vendor folder already has all the required modules

##Shopping Cart Feature (by Jeremias Mendoza)

This feature allows for users to add books to their shopping cart, retrieve the subtotal of their cart, retrieve the list of books in their cart, and remove books from their cart.

### API Endpoints

All endpoints are accessed through the base URL: http://localhost:8080/v1

--- 

### Add a book to Shopping Cart

**POST** `/shopping_cart_books`

JSON Body:
{
    "user_id": "{user_id}",
    "book_id": "{book_id}"
}

Adds a specific book to a user's shopping cart.

**Example:** http://localhost:8080/v1/shopping_cart_books
{
    "user_id": "80df9960-5f09-468f-a365-f2f404d01b0",
    "book_id": "cf5b0bf7-2617-44c1-b864-cc7f73782f09"
}

---

###Retrieve Shopping Cart Subtotal


