package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// creates a new wishlist for the given user_id
func (apiCgf *apiConfig) handlerCreateWishlist(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserID   uuid.UUID `json:"user_id"`
		ListName string    `json:"list_name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Checking the number of wishlists that the user has
	count_wishlists, err := apiCgf.DB.CountUserWishlists(r.Context(), params.UserID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to check wishlist count for '%v'", params.UserID))
		return
	}
	// Do not create wishlist if the user has previously created 3 wishlists.
	if count_wishlists >= 3 {
		respondWithError(w, 400, fmt.Sprintf("The user with id '%v' already has 3 wishlists", params.UserID))
		return
	}

	// Create the wishlist if no errors occur
	wishlist, err := apiCgf.DB.CreateWishlist(r.Context(), database.CreateWishlistParams{
		ID:       uuid.New(),
		UserID:   params.UserID,
		ListName: params.ListName,
	})

	if err != nil {
		// Check the user with the given id exists in the database
		if strings.Contains(err.Error(), "foreign key constraint") {
			respondWithError(w, 400, fmt.Sprintf("User with id '%v' does not exist", params.UserID))
		}
		// Check the names of each wishlist is unique for the same user id
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(w, 400, "Coudln't create the wishlist: Duplicate names are not allowed")
		}
		return
	}

	responseWithJSON(w, 201, databaseWishlistToWishlist(wishlist))
}

// add book to a wishlist
func (apiCgf *apiConfig) handlerAddBookToWishlist(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		WishlistID uuid.UUID `json:"wishlist_id"`
		BookID     uuid.UUID `json:"book_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Add book to wishlist
	book_to_wishlist, err := apiCgf.DB.AddBookToWishlist(r.Context(), database.AddBookToWishlistParams{
		ID:         uuid.New(),
		WishlistID: params.WishlistID,
		BookID:     params.BookID,
	})

	if err != nil {
		// Check the wishlist with the given id exists in the database
		if strings.Contains(err.Error(), "foreign key constraint") {
			respondWithError(w, 400, fmt.Sprintf("Wishlist with id '%v' does not exist", params.WishlistID))
		}
		// Check if a duplicate is trying to be added
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(w, 400, "Coudln't add book to wishlist: Duplicate books are not allowed")
		}
		return
	}

	responseWithJSON(w, 201, databaseBookWithWishlistToBookWithWishlist(book_to_wishlist))
}

// remove book from wishlist
func (apiCfg *apiConfig) handlerRemoveBookFromWishlist(w http.ResponseWriter, r *http.Request) {
	// parsing the parameters
	type parameters struct {
		WishlistID     uuid.UUID `json:"wishlist_id"`
		ToShoppingCart string    `json:"to_shopping_cart"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// setting book ID to be read from url
	wishlistBookIDStr := chi.URLParam(r, "wishlistBookID")
	wishlistBookID, err := uuid.Parse(wishlistBookIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse wishlist book id: %v", err))
	}

	// Deleting book from the wishlist
	err = apiCfg.DB.DeleteBookFromWishlist(r.Context(), database.DeleteBookFromWishlistParams{
		WishlistID: params.WishlistID,
		BookID:     wishlistBookID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete book from wishlist: %v", err))
		return
	}

	responseWithJSON(w, 200, fmt.Sprintf(
		"Book was successfully removed from the wishlist with ID: %v", params.WishlistID,
	))

	// Check if the book should be added to the user's shopping cart
	if params.ToShoppingCart == "yes" {
		// Get the user ID from wishlist ID
		userID, err := apiCfg.DB.GetUserIDByWishlistID(r.Context(), params.WishlistID)
		if err != nil {
			respondWithError(w, 500, "Failed to get user ID for wishlist")
			return
		}

		// Get the user's shopping cart
		cartID, err := apiCfg.DB.GetShoppingCartByUserID(r.Context(), userID)
		if err != nil {
			respondWithError(w, 500, "Failed to get user's shopping cart")
			return
		}

		// Add book to cart
		err = apiCfg.DB.AddBookToCart(r.Context(), database.AddBookToCartParams{
			ID:       uuid.New(),
			CartID:   cartID,
			BookID:   wishlistBookID,
			Quantity: sql.NullInt32{Int32: 1, Valid: true},
		})
		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("Error adding book to cart: %v", err))
			return
		}
		responseWithJSON(w, 200, fmt.Sprintf(
			"Book was successfully added to shopping cart with user ID: %v", userID,
		))
	}
}

// list all books from wishlist
