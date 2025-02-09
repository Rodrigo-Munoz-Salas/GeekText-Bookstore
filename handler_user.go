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

	user, err := apiCgf.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:           uuid.New(),
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

	responseWithJSON(w, 201, databaseUserToUser(user))
}
