package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/google/uuid"
)

// helper function to handle null strings
func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

// creates a new user
func (apiCgf *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User created (stub function)")
	type parameters struct {
		Username      string `json:"username"`
		Password_hash string `json:"password_hash"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		Home_address  string `json:"home_address"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Generate a new user ID
	userID := uuid.New()

	// Insert user into the database
	user, err := apiCgf.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:           userID,
		Username:     params.Username,
		PasswordHash: params.Password_hash,
		Name:         toNullString(params.Name),
		Email:        toNullString(params.Email),
		HomeAddress:  toNullString(params.Home_address),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Coudln't create user: %v", err))
		return
	}

	// Insert shopping cart for the new user
	_, err = apiCgf.DB.CreateShoppingCart(r.Context(), database.CreateShoppingCartParams{
		ID:     uuid.New(),
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("User created, but failed to create shopping cart: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseUserToUser(user))
}

// Gets user by username

// creates a new user
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	
	type parameters struct {

		Username      string `json:"username"`

	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByUsername(r.Context(), params.Username)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("User not found: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(user))


}