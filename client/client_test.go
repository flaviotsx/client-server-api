package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestGetQuotation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Quotation{Bid: 5.23}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", ts.URL+"/cotacao", nil)
	if err != nil {
		t.Errorf("Erro ao criar uma nova requisição %v", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Erro ao enviar a resposta %v", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Erro ao ler a resposta %v", err)
		return
	}

	var quotation Quotation
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		t.Errorf("Erro ao deserializar a resposta %v", err)
		return
	}

	if quotation.Bid != 5.23 {
		t.Errorf("Erro ao deserializar a resposta %v", err)
		return
	}

	content, err := os.ReadFile("cotacao.txt")
	if err != nil {
		t.Errorf("Erro ao ler a resposta %v", err)
		return
	}

	if string(content) != string(body) {
		t.Errorf("Erro ao escrever a resposta %v", err)
		return
	}

	if !bytes.Equal(content, body) {
		t.Errorf("Erro ao escrever a resposta %v", err)
		return
	}
}
