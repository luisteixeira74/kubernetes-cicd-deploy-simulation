package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	RootHandler(w, req) // Usando RootHandler corretamente

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
	}
}

func TestURLNotFound(t *testing.T) {
	// Simula um servidor com apenas a rota /hello registrada
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", RootHandler) // Usando RootHandler corretamente

	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("esperado status 404, obteve %d", resp.StatusCode)
	}
}
