package main

import (
	"database/sql"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/google/uuid"
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
