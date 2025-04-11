package main

import (
	"database/sql"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/google/uuid"

	"time"
)

type User struct {
	ID           uuid.UUID      `json:"id"`
	Username     string         `json:"username"`
	PasswordHash string         `json:"password_hash"`
	Name         sql.NullString `json:"name"`
	Email        sql.NullString `json:"email"`
	HomeAddress  sql.NullString `json:"home_address"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:           dbUser.ID,
		Username:     dbUser.Username,
		PasswordHash: dbUser.PasswordHash,
		Name:         dbUser.Name,
		Email:        dbUser.Email,
		HomeAddress:  dbUser.HomeAddress,
	}
}

type CreditCard struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	CardNumber     string    `json:"card_number"`
	ExpirationDate time.Time `json:"expiration_date"`
	Cvv            string    `json:"cvv"`
}

func databaseCardToCard(dbCard database.CreditCard) CreditCard {

	return CreditCard{
		ID:             dbCard.ID,
		UserID:         dbCard.UserID,
		CardNumber:     dbCard.CardNumber,
		ExpirationDate: dbCard.ExpirationDate,
		Cvv:            dbCard.Cvv,
	}
}

type Wishlist struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	ListName string    `json:"list_name"`
}

func databaseWishlistToWishlist(dbWishlist database.Wishlist) Wishlist {
	return Wishlist{
		ID:       dbWishlist.ID,
		UserID:   dbWishlist.UserID,
		ListName: dbWishlist.ListName,
	}
}

type WishlistBook struct {
	ID         uuid.UUID `json:"id"`
	WishlistID uuid.UUID `json:"wishlist_id"`
	BookID     uuid.UUID `json:"book"`
}

func databaseBookWithWishlistToBookWithWishlist(dbWishlistBooks database.WishlistBook) WishlistBook {
	return WishlistBook{
		ID:         dbWishlistBooks.ID,
		WishlistID: dbWishlistBooks.WishlistID,
		BookID:     dbWishlistBooks.BookID,
	}
}

type Book struct {
	ID            uuid.UUID      `json:"id"`
	Isbn          string         `json:"isbn"`
	Title         string         `json:"title"`
	Description   sql.NullString `json:"description"`
	Price         string         `json:"price"`
	Genre         string         `json:"genre"`
	PublisherID   uuid.NullUUID  `json:"publisher_id"`
	YearPublished int32          `json:"year_published"`
	CopiesSold    int32          `json:"copies_sold"`
	Author        string         `json:"author"`
}

func databaseBookToBook(dbBook database.Book) Book {
	return Book{
		ID:            dbBook.ID,
		Isbn:          dbBook.Isbn,
		Title:         dbBook.Title,
		Description:   dbBook.Description,
		Price:         dbBook.Price,
		Genre:         dbBook.Genre,
		PublisherID:   dbBook.PublisherID,
		YearPublished: dbBook.YearPublished,
		CopiesSold:    dbBook.CopiesSold,
		Author:        dbBook.Author,
	}
}

func databaseBooksToBooks(dbBooks []database.Book) []Book {
	books := []Book{}
	for _, dbBook := range dbBooks {
		books = append(books, databaseBookToBook(dbBook))
	}
	return books
}
