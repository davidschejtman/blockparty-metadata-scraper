// Pode ser em main.go ou em um novo arquivo, por exemplo, fetcher.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// Outros imports necessários
)

func fetchMetadata(cid string) (*Metadata, error) {
	url := fmt.Sprintf("https://ipfs.io/ipfs/%s", cid)
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
