package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"fmt"
	

	"github.com/Rodrigo-Munoz-Salas/GeekText-Bookstore/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

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

	// import our database connection from .env file
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// Setting database connectivity with db url of local host
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	// Setting database api configuration between app and postgres
	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

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
	//v1Router.Post() <- this is to how create the router

	// START FEATURE IMPLEMENTATIONS
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUser)

	v1Router.Post("/book_admin", apiCfg.handlerCreateBook)

	v1Router.Post("/wishlists", apiCfg.handlerCreateWishlist)
	v1Router.Post("/wishlist_books", apiCfg.handlerAddBookToWishlist)
	v1Router.Delete("/wishlist_books/{wishlistBookID}", apiCfg.handlerRemoveBookFromWishlist)
	v1Router.Get("/wishlist_books", apiCfg.handlerGetWishlistBooks)

	v1Router.Post("/shopping_cart_books", apiCfg.handlerAddBookToCart)
	v1Router.Get("/shopping_cart_books/subtotal", apiCfg.handlerGetCartSubtotal)
	v1Router.Get("/shopping_cart_books/list", apiCfg.handlerGetCartBooks)
	v1Router.Delete("/shopping_cart_books/delete", apiCfg.handlerDeleteBookFromCart)
	v1Router.Post("/rating", apiCfg.handlerPostRating)
	v1Router.Post("/comments", apiCfg.handlerPostComment)
	v1Router.Get("/rating/{bookID}", apiCfg.handlerAvgRating)
	v1Router.Get("/comments/{bookID}", apiCfg.handlerGetComments)

	// STOP FEATURE IMPLEMENTATIONS, DO NOT TOUCH BELOW

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
	log.Println("Database connectivity with Postgres was succesfully done!")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
