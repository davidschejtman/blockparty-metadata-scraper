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
	// Carregar as variáveis de ambiente
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.DB.Close()

	// Configuração do roteador
	r := mux.NewRouter()
	r.HandleFunc("/tokens", api.GetTokens).Methods("GET")
	r.HandleFunc("/tokens/{cid}", api.GetTokenByCID).Methods("GET")

	// Definindo a porta
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // Porta padrão, caso não especificada
	}

	// Iniciando o servidor
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
