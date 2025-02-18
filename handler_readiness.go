package main

import (
	"net/http"
)

// function signature to define an http handler as the go
// standard library expects
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, map[string]string{
		"confirm_message": "Health Point",
	})
}
