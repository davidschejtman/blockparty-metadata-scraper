package main

import (
	"io/ioutil"
	"log"
	"strings"

	"blockparty-metadata-scraper/internal/db"
	"blockparty-metadata-scraper/internal/scraper"
)

func main() {
	// Initialize the connection to the database
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.DB.Close()

	// Read CIDs from a file
	data, err := ioutil.ReadFile("files/ipfs_cids.csv")
	if err != nil {
		log.Fatalf("Failed to read CIDs file: %v", err)
	}

	// Converting the file contents into a CID slice
	cids := strings.Split(string(data), "\n")

	// Run the scraping and storage process for each CID
	scraper.FetchAndStoreMetadata(cids)
}
