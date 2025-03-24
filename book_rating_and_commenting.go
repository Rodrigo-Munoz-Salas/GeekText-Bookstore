package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

//Post rating function
func (apiCfg *apiConfig) handlerPostRating(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		BookID uuid.UUID `json:"book_id"`
		Rating int32     `json:"rating"`
		UserID uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	rating, err := apiCfg.DB.CreateRating(r.Context(), database.CreateRatingParams{
		ID: 	uuid.New(),
		BookID: params.BookID,
		Rating: params.Rating,
		UserID: params.UserID,
	})

	if (err != nil) {
		respondWithError(w, 400, fmt.Sprintf("Failed to create new rating: %v", err))
		return
	}
	responseWithJSON(w, 201, rating)
} 

//Post comment function
func (apiCfg *apiConfig) handlerPostComment(w http.ResponseWriter, r *http.Request) {	decoder := json.NewDecoder(r.Body)
	type parameters struct {
		BookID uuid.UUID `json:"book_id"`
		Comment string   `json:"comment"`
		UserID uuid.UUID `json:"user_id"`
	}
	decoder = json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	comment, err := apiCfg.DB.CreateComment(r.Context(), database.CreateCommentParams{
		ID: 	 uuid.New(),
		BookID:  params.BookID,
		Comment: params.Comment,
		UserID:  params.UserID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to create comemnt: %v", err))
		return
	}
	responseWithJSON(w, 201, comment)
}

//Gets average rating for each book
func (apiCfg *apiConfig) handlerAvgRating(w http.ResponseWriter, r *http.Request) {
	bookIDStr := chi.URLParam(r, "bookID")
	bookID, err := uuid.Parse(bookIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid Book ID: %v", err))
		return
	}
	avgRatingInterface, err := apiCfg.DB.GetAveRatingByBook(r.Context(), bookID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to get average rating: %v", err))
		return
	}
	if avgRatingInterface == nil {
		responseWithJSON(w, 200, map[string]float64{"average_rating": 0})
		return
	}
	var avgRating float64
	switch v := avgRatingInterface.(type) {
		case float64:
			avgRating = v
		case int64:
			avgRating = float64(v)
		case string:
			parsed, err := strconv.ParseFloat(v, 64)
			if err != nil {
				respondWithError(w, 500, fmt.Sprintf("Unexpected string type for average rating: %v", err))
				return 
			}
			avgRating = parsed
		case []uint8:
			parsed, err := strconv.ParseFloat(string(v), 64)
			if err != nil {
				respondWithError(w, 500, fmt.Sprintf("Unexpected []uint8 type for average rating: %v", err))
				return
			}
			avgRating = parsed
		default:
			respondWithError(w, 500, fmt.Sprintf("Unexpected type for average rating: %T", avgRatingInterface))
			return
	}
	responseWithJSON(w, 200, map[string]float64{"average_rating": avgRating})
}

//Gets all comments for a book
func (apiCfg *apiConfig) handlerGetComments(w http.ResponseWriter, r *http.Request) {
	bookIDStr := chi.URLParam(r, "bookID")
	bookID, err := uuid.Parse(bookIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid Book ID: %v", err))
		return
	}
	comments, err := apiCfg.DB.GetCommentsByBook(r.Context(), bookID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to get comments: %v", err))
		return
	}
	responseWithJSON(w, 200, comments)
}