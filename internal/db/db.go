package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	// Carrega as variáveis de ambiente diretamente ou use uma biblioteca como godotenv para carregá-las de um arquivo .env
	err := godotenv.Load() // Procura por .env no diretório atual
	if err != nil {
		log.Println("No .env file found or error loading it, falling back to system environment variables.")
	}
	/*
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		// Constrói a string de conexão
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)*/

	connStr := fmt.Sprintf("host=candidate-testing.co2sjmg0hdpm.us-east-2.rds.amazonaws.com port=5432 user=dschejtman password=xj98Zs0f7sl2idk3ls dbname=dschejtman sslmode=require")
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Verifica a conexão com o banco de dados
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	log.Println("Connected to the database successfully.")

	// Cria a tabela metadata se ela não existir
	if err = createMetadataTable(); err != nil {
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
