package main

import (
	"log"
	"net/http"
	"os"

	"blockparty-metadata-scraper/internal/api"
	"blockparty-metadata-scraper/internal/db"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.DB.Close()

	// Router configuration
	r := mux.NewRouter()
	r.HandleFunc("/tokens", api.GetTokens).Methods("GET")
	r.HandleFunc("/tokens/{cid}", api.GetTokenByCID).Methods("GET")

	// Defining the port
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Starting the server
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
