package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchMetadataSuccess(t *testing.T) {
	// Simular resposta da API IPFS
	mockResponse := Metadata{
		Image:       "http://example.com/image.png",
		Description: "A test image",
		Name:        "TestImage",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Simula status 200 OK
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Substituir baseURL pelo URL do servidor mock
	oldBaseURL := baseURL
	baseURL = server.URL + "/"
	defer func() { baseURL = oldBaseURL }() // Restaurar após o teste

	// Executar fetchMetadata com um CID mock
	metadata, err := fetchMetadata("QmTestCID")
	if err != nil {
		t.Fatalf("fetchMetadata failed: %v", err)
	}

	// Verificar se os metadados retornados correspondem à simulação
	if metadata.Image != mockResponse.Image || metadata.Description != mockResponse.Description || metadata.Name != mockResponse.Name {
		t.Errorf("fetchMetadata returned wrong metadata, got %+v, want %+v", metadata, mockResponse)
	}
}

func TestFetchMetadataNon200Response(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Simula status 404 Not Found
	}))
	defer server.Close()

	// Substituir baseURL pelo URL do servidor mock
	oldBaseURL := baseURL
	baseURL = server.URL + "/"
	defer func() { baseURL = oldBaseURL }() // Restaurar após o teste

	// Executar fetchMetadata esperando um erro
	_, err := fetchMetadata("QmInvalidCID")
	if err == nil {
		t.Fatal("fetchMetadata expected to fail for non-200 response, but it did not")
	}

	// Aqui, você também pode querer verificar a mensagem de erro específica, se necessário
}
