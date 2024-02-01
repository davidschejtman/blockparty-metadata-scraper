// cidreader.go
package main

import (
	"encoding/csv"
	"io"
	"os"
)

// ReadCIDsFromFile lê um arquivo CSV e retorna uma lista de CIDs.
func ReadCIDsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var cids []string

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		cids = append(cids, record[0])
	}

	return cids, nil
}
