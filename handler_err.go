package main

import "net/http"

// helper function to respond with error in json format
// function signature for HTTP requests in go. Do not change
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
