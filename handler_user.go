package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/google/uuid"

	"strconv"
	"strings"
	"time"
)

// Helper function to handle null strings
func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

// Creates a new user
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

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Username string `json:"username"`
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

// Updates user details given username and updated feilds
func (apiCfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Username      string `json:"username"`
		Password_hash string `json:"password_hash,omitempty"`
		Name          string `json:"name,omitempty"`
		Home_address  string `json:"home_address,omitempty"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	_, err = apiCfg.DB.GetUserByUsername(r.Context(), params.Username)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("User not found: %v", err))
		return
	}

	updateParams := database.UpdateUserByUsernameParams{
		Username: params.Username,
	}

	if params.Password_hash != "" {
		updateParams.Column2 = params.Password_hash
	}

	if params.Name != "" {
		updateParams.Column3 = params.Name
	}

	if params.Home_address != "" {
		updateParams.Column4 = params.Home_address
	}

	_, err = apiCfg.DB.UpdateUserByUsername(r.Context(), updateParams)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("User cannot be updated: %v", err))
		return
	}

	newUser, err := apiCfg.DB.GetUserByUsername(r.Context(), params.Username)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("User not found: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(newUser))

}

//creates User's credit card

func (apiCfg *apiConfig) handlerUserCreditCard(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Username       string `json:"username"`
		CardNumber     string `json:"card_number"`
		ExpirationDate string `json:"expiration_date"`
		Cvv            string `json:"cvv"`
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

	creditCardID := uuid.New()

	expirationDate, err := parseExpirationDate(params.ExpirationDate)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid expiration date format: %v", err))
		return
	}

	creditCard, err := apiCfg.DB.CreateUserCreditCard(r.Context(), database.CreateUserCreditCardParams{

		ID:             creditCardID,
		UserID:         user.ID,
		CardNumber:     params.CardNumber,
		ExpirationDate: expirationDate,
		Cvv:            params.Cvv,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Credit Card not created: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseCardToCard(creditCard))

}

// Helper function to convert "MM/YYYY" to "YYYY-MM-DD" (last day of the month)
func parseExpirationDate(expiration string) (time.Time, error) {
	parts := strings.Split(expiration, "/")
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("expected format MM/YYYY")
	}

	month, err := strconv.Atoi(parts[0])
	if err != nil || month < 1 || month > 12 {
		return time.Time{}, fmt.Errorf("invalid month")
	}

	year, err := strconv.Atoi(parts[1])
	if err != nil || year < time.Now().Year() {
		return time.Time{}, fmt.Errorf("invalid year")
	}

	// Get the last day of the given month & year
	lastDay := lastDayOfMonth(year, time.Month(month))

	return time.Date(year, time.Month(month), lastDay, 0, 0, 0, 0, time.UTC), nil
}

// Helper function to get last day of a given month
func lastDayOfMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
