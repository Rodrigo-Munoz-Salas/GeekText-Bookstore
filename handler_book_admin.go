package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/google/uuid"
)

// creates a new book
func (apiCfg *apiConfig) handlerCreateBook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ISBN          string  `json:"isbn"`
		Title         string  `json:"title"`
		Description   string  `json:"description"`
		Price         float64 `json:"price"`
		Genre         string  `json:"genre"`
		PublisherName string  `json:"publisher_name"`
		YearPublished int     `json:"year_published"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Check if publisher is already in the database
	publisherID, err := apiCfg.DB.GetPublisherByName(r.Context(), params.PublisherName)
	if err != nil {
		if err == sql.ErrNoRows {
			// Add publisher to the DB
			newPublisherID := uuid.New()
			_, err = apiCfg.DB.CreatePublisher(r.Context(), database.CreatePublisherParams{
				ID:   newPublisherID,
				Name: params.PublisherName,
			})
			if err != nil {
				respondWithError(w, 500, fmt.Sprintf("Could not create publisher: %v", err))
				return
			}
			publisherID = newPublisherID
		} else {
			respondWithError(w, 500, fmt.Sprintf("Database error: %v", err))
			return
		}
	}

	// Add new Book to the Database
	book, err := apiCfg.DB.CreateBook(r.Context(), database.CreateBookParams{
		ID:            uuid.New(),
		Isbn:          params.ISBN,
		Title:         params.Title,
		Description:   toNullString(params.Description),
		Price:         fmt.Sprintf("%.2f", params.Price),
		Genre:         params.Genre,
		PublisherID:   uuid.NullUUID{UUID: publisherID, Valid: true},
		YearPublished: int32(params.YearPublished),
	})

	// Check if there is an error while adding the book to the DB
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not create the book: %v", err))
		return
	}

	// Respond with success message
	responseWithJSON(w, 200, databaseBookToBook(book))
}
