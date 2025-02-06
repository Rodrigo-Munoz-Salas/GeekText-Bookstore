package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Setting port to run the GeekText App
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// CREATE A .env FILE AND ADD PORT={YOUR_PORT}
	// Uncomment this line and test it
	// fmt.Printf("PORT is: %v", portString)

	// Creating the router to bind the endpoints
	router := chi.NewRouter()

	// Make requests to the server from a browser
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Hook up handlers to specific http methods and paths
	v1Router := chi.NewRouter()

	// Checking health endpoint and error handler
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	// Mounting router with v1 router
	router.Mount("/v1", v1Router)

	// Connecting router to an http server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	// Server starts running here, handleling HTTP requests
	// If an error ocurred, the server will inmediately stop and log the error
	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
