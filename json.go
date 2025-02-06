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
	type errResponse struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal the payload into a json string
	dat, err := json.Marshal(payload)

	// internal server error if Marshal process fails
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	// write the response header and body
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
