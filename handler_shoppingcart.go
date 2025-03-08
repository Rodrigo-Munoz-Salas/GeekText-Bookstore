package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/google/uuid"
)

// adds a book to the shopping cart
func (apiCfg *apiConfig) handlerAddBookToCart(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserID uuid.UUID `json:"user_id"`
		BookID uuid.UUID `json:"book_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Add book to cart
	cartItem, err := apiCfg.DB.AddBookToCart(r.Context(), database.AddBookToCartParams{
		ID:     uuid.New(),
		UserID: params.UserID,
		BookID: params.BookID,
	})

	if err != nil {
		if strings.Contains(err.Error(), "foreign key constraint") {
			respondWithError(w, 400, fmt.Sprintf("User with id '%v' does not exist", params.UserID))
		}
		return
	}

	responseWithJSON(w, 201, databaseCartItemToCartItem(cartItem))
}

func databaseCartItemToCartItem(cartItem database.CartItem) interface{} {
	panic("unimplemented")
}

// retrieves the subtotal price of all items in the user's cart
func (apiCfg *apiConfig) handlerGetCartSubtotal(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserID uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	subtotal, err := apiCfg.DB.GetCartSubtotal(r.Context(), params.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to retrieve cart subtotal")
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]float64{"subtotal": subtotal})
}
