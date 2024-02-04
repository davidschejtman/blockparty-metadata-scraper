package models

// Metadata representa a estrutura dos metadados associados a um CID IPFS.
type Metadata struct {
	CID         string `json:"cid"`         // Identificador único do conteúdo no IPFS
	Image       string `json:"image"`       // URL da imagem associada ao metadado
	Description string `json:"description"` // Descrição do conteúdo
	Name        string `json:"name"`        // Nome do conteúdo
}
