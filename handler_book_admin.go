package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

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
		CopiesSold    int     `json:"copies_sold"`
		AuthorName    string  `json:"author"`
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

	// Split AuthorName into first and last name
	nameParts := strings.Split(params.AuthorName, " ")
	if len(nameParts) != 2 {
		respondWithError(w, 400, "Author name must contain both first and last names")
		return
	}

	firstName := nameParts[0]
	lastName := nameParts[1]

	// Check if author is already in the database using the GetAuthorByNameParams struct
	authorParams := database.GetAuthorByNameParams{
		FirstName: firstName,
		LastName:  lastName,
	}

	_, err = apiCfg.DB.GetAuthorByName(r.Context(), authorParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Author was not found
			respondWithError(w, 404, fmt.Sprintf("Author not found: %v", err))
			return
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
		CopiesSold:    int32(params.CopiesSold),
		Author:        params.AuthorName,
	})

	// Check if there is an error while adding the book to the DB
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not create the book: %v", err))
		return
	}

	// Respond with success message
	responseWithJSON(w, 200, databaseBookToBook(book))
}

func (apiCfg *apiConfig) handlerGetBookByIsbn(w http.ResponseWriter, r *http.Request) {

	// Define the parameters structure to receive the ISBN
	type parameters struct {
		ISBN string `json:"isbn"`
	}

	// Parse the request body
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Fetch the book details using the ISBN from the database
	book, err := apiCfg.DB.GetBookByISBN(r.Context(), params.ISBN)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Book not found: %v", err))
		return
	}

	// Respond with the book details as JSON
	responseWithJSON(w, 200, book)
}

func (apiCfg *apiConfig) handlerCreateAuthor(w http.ResponseWriter, r *http.Request) {
	// Define the structure to map incoming JSON to
	type parameters struct {
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		Biography   string    `json:"biography"`
		PublisherID uuid.UUID `json:"publisher_id"`
	}

	// Parse the incoming request body into the parameters struct
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Handle the biography as sql.NullString (if empty, it will be null)
	biography := sql.NullString{}
	if params.Biography != "" {
		biography = sql.NullString{String: params.Biography, Valid: true}
	}

	// Handle the publisher_id as uuid.NullUUID
	publisherID := uuid.NullUUID{UUID: params.PublisherID, Valid: true}

	// Create the author in the database
	authorID, err := apiCfg.DB.CreateAuthor(r.Context(), database.CreateAuthorParams{
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		Biography:   biography,
		PublisherID: publisherID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating author: %v", err))
		return
	}

	// Respond with the ID of the created author (or you can fetch and return more details if needed)
	// Here, just returning the ID as a simple response
	responseWithJSON(w, 200, map[string]interface{}{
		"message": "Author successfully created.",
		"id":      authorID,
	})
}

func (apiCfg *apiConfig) handlerGetBooksByAuthorId(w http.ResponseWriter, r *http.Request) {
	// Define a struct to hold the request body
	var requestBody struct {
		AuthorID uuid.UUID `json:"author_id"`
	}

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		// Send back a 400 Bad Request response with an error message
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	// Get book IDs by author ID from the database
	bookIds, err := apiCfg.DB.GetBookIdsByAuthorId(r.Context(), requestBody.AuthorID)
	if err != nil {
		// Send back a 500 Internal Server Error response if there's an issue fetching book IDs
		responseWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error retrieving book IDs"})
		return
	}

	// Fetch book details for each book ID
	var books []database.Book
	for _, bookId := range bookIds {
		book, err := apiCfg.DB.GetBookDetailsByBookId(r.Context(), bookId)
		if err != nil {
			// Handle error for individual book fetching (you may log or ignore it)
			continue // Skip this book if there's an error
		}
		books = append(books, book)
	}

	// Send back a 200 OK response with the books as JSON
	responseWithJSON(w, 200, books)
}
