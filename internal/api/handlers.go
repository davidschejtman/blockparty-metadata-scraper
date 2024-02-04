package api

import (
	"blockparty-metadata-scraper/internal/db"
	"blockparty-metadata-scraper/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetTokens é o handler para o endpoint GET /tokens. Ele busca todos os metadados armazenados e os retorna.
func GetTokens(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var metadatas []models.Metadata
	query := `SELECT cid, image, description, name FROM metadata;`

	rows, err := db.DB.Query(query)
	if err != nil {
		http.Error(w, "Error fetching metadata from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Metadata
		if err := rows.Scan(&m.CID, &m.Image, &m.Description, &m.Name); err != nil {
			http.Error(w, "Error reading metadata rows", http.StatusInternalServerError)
			return
		}
		metadatas = append(metadatas, m)
	}

	if err := json.NewEncoder(w).Encode(metadatas); err != nil {
		http.Error(w, "Error encoding response into JSON", http.StatusInternalServerError)
	}
}

// GetTokenByCID é o handler para o endpoint GET /tokens/{cid}. Ele busca metadados específicos por CID.
func GetTokenByCID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	cid := params["cid"]

	var metadata models.Metadata
	query := `SELECT cid, image, description, name FROM metadata WHERE cid = $1;`

	err := db.DB.QueryRow(query, cid).Scan(&metadata.CID, &metadata.Image, &metadata.Description, &metadata.Name)
	if err != nil {
		http.Error(w, "Metadata not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(metadata); err != nil {
		http.Error(w, "Error encoding response into JSON", http.StatusInternalServerError)
	}
}
