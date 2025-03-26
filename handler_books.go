package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// Retrieve books by genre
func (apiCfg *apiConfig) HandlerGetBooksByGenre(w http.ResponseWriter, r *http.Request) {
	genre := chi.URLParam(r, "genre")

	books, err := apiCfg.DB.GetBooksByGenre(r.Context(), genre)
	if err != nil {
		log.Printf("Database error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch books")
		return
	}

	responseWithJSON(w, http.StatusOK, books)
}

// Retrieve top 10 best-selling books
func (apiCfg *apiConfig) HandlerGetTopSellers(w http.ResponseWriter, r *http.Request) {
	books, err := apiCfg.DB.GetTopSellingBooks(r.Context())
	if err != nil {
		log.Printf("Database error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch top sellers")
		return
	}

	responseWithJSON(w, http.StatusOK, books)
}

// Retrieve books by rating
func (apiCfg *apiConfig) HandlerGetBooksByRating(w http.ResponseWriter, r *http.Request) {
	ratingStr := chi.URLParam(r, "rating")
	rating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid rating parameter")
		return
	}

	books, err := apiCfg.DB.GetBooksByRating(r.Context(), int32(rating))
	if err != nil {
		log.Printf("Database error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch books by rating")
		return
	}

	responseWithJSON(w, http.StatusOK, books)
}

// Apply discount to books from a specific publisher
func (apiCfg *apiConfig) HandlerApplyDiscountToPublisher(w http.ResponseWriter, r *http.Request) {
	discountStr := r.URL.Query().Get("discount")
	publisherIDStr := r.URL.Query().Get("publisher_id")

	discount, err := strconv.ParseFloat(discountStr, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid discount value")
		return
	}

	publisherID, err := uuid.Parse(publisherIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid publisher ID")
		return
	}

	params := database.ApplyDiscountToPublisherParams{
		DiscountPercent: discount,
		PublisherID:     publisherID,
	}

	log.Printf("Applying %.2f%% discount to publisher: %s", discount*100, publisherID)

	err = apiCfg.DB.ApplyDiscountToPublisher(r.Context(), params)
	if err != nil {
		log.Printf("Database error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to apply discount")
		return
	}

	log.Printf("Discount applied successfully to publisher: %s", publisherID)
	responseWithJSON(w, http.StatusOK, map[string]string{
		"message": "Discount applied successfully",
	})
}
