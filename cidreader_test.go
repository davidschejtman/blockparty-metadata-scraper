package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadCIDsFromFile(t *testing.T) {
	// Criar um arquivo CSV temporário
	content := "bafkreifwbyviygstiqmjcijju33r6scctuyxciqiepcrs2ym2bbpf7c7rq\r\nafkreifovhtvvrx5jmo2b4ne2hoyk4t3c276jc7weva5s57ilupiiuqg2y"
	tmpFile, err := ioutil.TempFile("", "test_cids_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Limpa o arquivo temporário após o teste

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Testar ReadCIDsFromFile
	cids, err := ReadCIDsFromFile(tmpFile.Name())
	if err != nil {
		t.Errorf("ReadCIDsFromFile returned an error: %v", err)
	}
	if len(cids) != 2 {
		t.Errorf("Expected 2 CIDs, got %d", len(cids))
	}
}
