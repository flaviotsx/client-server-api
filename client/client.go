package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Quotation struct {
	Bid float64 `json:"bid"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Erro ao criar uma nova requisição", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar a resposta", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta", err)
	}

	var quotation *Quotation

	err = json.Unmarshal(body, &quotation)
	if err != nil {
		fmt.Println("Erro ao deserializar a resposta", err)
	}

	fmt.Printf("Dollar: %.2f\n", quotation.Bid)

	err = os.WriteFile("cotacao.txt", body, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever a resposta", err)
		return
	}

}
