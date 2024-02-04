package test

import (
	"blockparty-metadata-scraper/internal/api"
	"blockparty-metadata-scraper/internal/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetTokens(t *testing.T) {
	// Inicialização: substitua isso com a inicialização do seu banco de dados de teste
	db.InitDB()

	req, err := http.NewRequest("GET", "/tokens", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetTokens)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Aqui você pode adicionar mais verificações, como checar o corpo da resposta
}

func TestGetTokenByCID(t *testing.T) {
	// Inicialização: substitua isso com a inicialização do seu banco de dados de teste
	db.InitDB()

	req, err := http.NewRequest("GET", "/tokens/{cid}", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"cid": "bafkreifovhtvvrx5jmo2b4ne2hoyk4t3c276jc7weva5s57ilupiiuqg2y",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetTokenByCID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
