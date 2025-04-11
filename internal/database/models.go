// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	Biography   sql.NullString
	PublisherID uuid.NullUUID
}

type Book struct {
	ID            uuid.UUID
	Isbn          string
	Title         string
	Description   sql.NullString
	Price         string
	Genre         string
	PublisherID   uuid.NullUUID
	YearPublished int32
	CopiesSold    int32
	Author        string
}

type BookAuthor struct {
	BookID   uuid.UUID
	AuthorID uuid.UUID
}

type Comment struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	BookID    uuid.UUID
	Comment   string
	CreatedAt time.Time
}

type CreditCard struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	CardNumber     string
	ExpirationDate time.Time
	Cvv            string
}

type Publisher struct {
	ID   uuid.UUID
	Name string
}

type Rating struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	BookID    uuid.UUID
	Rating    int32
	CreatedAt time.Time
}

type ShoppingCart struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type ShoppingCartBook struct {
	ID       uuid.UUID
	CartID   uuid.UUID
	BookID   uuid.UUID
	Quantity sql.NullInt32
}

type User struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	Name         sql.NullString
	Email        sql.NullString
	HomeAddress  sql.NullString
}

type Wishlist struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	ListName string
}

type WishlistBook struct {
	ID         uuid.UUID
	WishlistID uuid.UUID
	BookID     uuid.UUID
}
