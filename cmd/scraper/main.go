package main

import (
	"io/ioutil"
	"log"
	"strings"

	"blockparty-metadata-scraper/internal/db"
	"blockparty-metadata-scraper/internal/scraper"
)

func main() {
	// Inicializa a conexão com o banco de dados
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.DB.Close()

	// Lê os CIDs a partir de um arquivo (substitua "path/to/your/ipfs_cids.csv" pelo caminho correto do arquivo)
	data, err := ioutil.ReadFile("files/ipfs_cids.csv")
	if err != nil {
		log.Fatalf("Failed to read CIDs file: %v", err)
	}

	// Convertendo o conteúdo do arquivo em uma slice de CIDs
	cids := strings.Split(string(data), "\n")

	// Executa o processo de raspagem e armazenamento para cada CID
	scraper.FetchAndStoreMetadata(cids)
}
