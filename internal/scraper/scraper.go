package scraper

import (
	"blockparty-metadata-scraper/internal/db"
	"blockparty-metadata-scraper/pkg/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// FetchAndStoreMetadata recebe uma lista de CIDs, busca seus metadados e os armazena no banco de dados.
func FetchAndStoreMetadata(cids []string) {
	for _, cid := range cids {
		metadata, err := FetchMetadata(cid)
		if err != nil {
			log.Printf("Erro ao buscar metadados para CID %s: %v\n", cid, err)
			continue
		}

		if err := StoreMetadata(metadata); err != nil {
			log.Printf("Erro ao armazenar metadados para CID %s: %v\n", cid, err)
			continue
		}
	}
}

// FetchMetadata faz uma requisição HTTP GET para buscar metadados de um CID específico.
func FetchMetadata(cid string) (*models.Metadata, error) {
	url := fmt.Sprintf("https://ipfs.io/ipfs/%s", cid)
	resp, err := http.Get(url)
	if err != nil {
		// Loga o erro da requisição
		log.Printf("Erro na requisição ao IPFS para CID %s: %v", cid, err)
		return nil, fmt.Errorf("erro na requisição ao IPFS: %v", err)
	}
	defer resp.Body.Close()

	// Verifica o código de status HTTP
	if resp.StatusCode != http.StatusOK {
		log.Printf("Erro: Código de status %d ao tentar acessar %s para CID %s", resp.StatusCode, url, cid)
		// Tratamento específico para códigos de status inesperados
		return nil, fmt.Errorf("status HTTP inesperado: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "application/json") {
		log.Printf("Tipo de conteúdo inesperado: %s para CID %s", contentType, cid)
		// Tratamento específico para tipos de conteúdo inesperados
		return nil, fmt.Errorf("tipo de conteúdo inesperado: %s", contentType)
	}

	// Processa o corpo da resposta
	var metadata models.Metadata
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		// Loga o erro de deserialização
		log.Printf("Erro ao deserializar a resposta para CID %s: %v", cid, err)
		return nil, fmt.Errorf("erro no unmarshal do JSON: %v", err)
	}

	return &metadata, nil
}

// StoreMetadata insere os metadados recuperados no banco de dados.
func StoreMetadata(metadata *models.Metadata) error {
	query := `INSERT INTO metadata (cid, image, description, name) VALUES ($1, $2, $3, $4) ON CONFLICT (cid) DO NOTHING;`
	_, err := db.DB.Exec(query, metadata.CID, metadata.Image, metadata.Description, metadata.Name)
	if err != nil {
		return fmt.Errorf("erro ao inserir metadados no banco de dados: %v", err)
	}

	log.Printf("Metadados inseridos com sucesso para CID %s\n", metadata.CID)
	return nil
}
