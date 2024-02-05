package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error // Declare err at the top of the function scope
	// Try to load environment variables from a .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or error loading it. Falling back to system environment variables.")
	}

	// Load environment variables or use default values
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Database environment variables are not set. Please check your .env file or system environment variables.")
	}

	// Builds the connection string based on environment variables
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	// Open the connection to the database
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Check the connection to the database
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	log.Println("Connected to the database successfully.")

	// Create the metadata table if it does not exist
	if err := createMetadataTable(); err != nil {
		return fmt.Errorf("error creating metadata table: %v", err)
	}

	return nil
}

func createMetadataTable() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS metadata (
		cid TEXT PRIMARY KEY,
		image TEXT,
		description TEXT,
		name TEXT
	);`

	if _, err := DB.Exec(createTableSQL); err != nil {
		return fmt.Errorf("error creating metadata table: %v", err)
	}

	log.Println("Metadata table verified/created successfully.")
	return nil
}
