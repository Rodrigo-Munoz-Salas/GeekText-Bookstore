package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerAddBookToCart(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserID uuid.UUID `json:"user_id"`
		BookID uuid.UUID `json:"book_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid user id: %v", err))
		return
	}

	// Retrieve the shopping cart for the user
	cartID, err := apiCfg.DB.GetShoppingCartByUserID(r.Context(), params.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, 404, "Shopping cart not found for this user")
			return
		}
		respondWithError(w, 500, fmt.Sprintf("Error retrieving shopping cart: %v", err))
		return
	}

	// Insert the book into the shopping cart or update the quantity
	// Default quantity of 1
	err = apiCfg.DB.AddBookToCart(r.Context(), database.AddBookToCartParams{
		ID:       uuid.New(),
		CartID:   cartID,
		BookID:   params.BookID,
		Quantity: sql.NullInt32{Int32: 1, Valid: true},
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error adding book to cart: %v", err))
		return
	}

	// Respond with success
	responseWithJSON(w, 200, "Book added to cart successfully")
}

// Retrieves Subtotal of cart by user_id
func (apiCfg *apiConfig) handlerGetCartSubtotal(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid user id: %v", err))
		return
	}

	// Retrieve the subtotal directly using user_id
	subtotal, err := apiCfg.DB.GetCartSubtotalByUserID(r.Context(), params.UserID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error retrieving cart subtotal: %v", err))
		return
	}

	// Respond with the subtotal
	responseWithJSON(w, 200, map[string]float64{"subtotal": subtotal})
}
