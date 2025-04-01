package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	// internal server error
	if code > 499 {
		log.Println("Responding with 5X error:", msg)
	}

	responseWithJSON(w, code, map[string]string{
		"error": msg,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}, message ...string) {
	// Default message if none is provided
	msg := ""
	if len(message) > 0 {
		msg = message[0] // Use the first provided message
	}

	/// Define a struct for the response
	type Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	// Response object
	response := Response{
		Message: msg,
		Data:    payload,
	}

	// Marshal the response into a JSON string
	dat, err := json.Marshal(response)

	// Internal server error if the Marshal process fails
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response header and body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
