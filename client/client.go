package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/murilocarbol/client-server-api/client/model"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	serverURL      = "http://localhost:8080/cotacao"
	requestTimeout = 3000000000 * time.Millisecond
	outputFile     = "cotacao.txt"
)

func BuscarCotacao() {

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		log.Panicf("erro a criar a requisição %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicf("Erro a realizar a requisição %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Panicf("o servidor retornou o status code: %v", resp.StatusCode)
	}

	var response *model.ServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Panicf("erro ao decodificar a resposta: %v", err)
	}

	content := fmt.Sprintf("Dólar: %s", response.Bid)
	if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
		log.Panicf("error ao gravar no arquivo")
	}

	log.Printf("arquivo com cotacao salvo com sucesso %s", outputFile)
}
