// fetcher.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// Outros imports necessários
)

// Define a URL base como uma variável global para permitir sua substituição nos testes
var baseURL = "https://ipfs.io/ipfs/"

func fetchMetadata(cid string) (*Metadata, error) {
	url := fmt.Sprintf("%s%s", baseURL, cid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request to IPFS: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	var metadata Metadata
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err := json.Unmarshal(body, &metadata); err != nil {
		return nil, fmt.Errorf("error unmarshaling metadata: %v", err)
	}

	return &metadata, nil
}
