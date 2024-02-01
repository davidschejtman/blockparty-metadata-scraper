package main

import (
	"fmt"
	"log"
)

func main() {
	cids, err := ReadCIDsFromFile("files\\ipfs_cids.csv")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo: %v", err)
	}

	for _, cid := range cids {
		metadata, err := fetchMetadata(cid)
		if err != nil {
			log.Printf("Erro ao buscar metadados para CID %s: %v", cid, err)
			continue
		}
		fmt.Printf("Metadados para CID %s: %+v\n", cid, metadata)
	}
}
