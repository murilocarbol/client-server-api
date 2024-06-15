package server

import (
	"context"
	"encoding/json"
	"github.com/murilocarbol/client-server-api/database"
	"github.com/murilocarbol/client-server-api/server/model"
	"github.com/murilocarbol/client-server-api/server/response"
	"github.com/murilocarbol/client-server-api/util"
	"log"
	"net/http"
	"time"
)

const (
	urlCotacao  = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	apiTimeout  = 200 * time.Millisecond
	serverPort  = ":8080"
	serverRoute = "/cotacao"
)

func SubirServer() {

	router := http.NewServeMux()
	router.HandleFunc(serverRoute, serverHandler)

	go util.SetupHealthCheck(router)

	http.ListenAndServe(serverPort, router)
}

func serverHandler(w http.ResponseWriter, r *http.Request) {

	// realiza a conexão com o sqlite
	database.UpDatabase()

	cotacao, err := capturarCotacao()
	if err != nil {
		log.Printf("erro ao obter cotacao: %v", err)
		http.Error(w, "nao foi possível obter a cotacao", http.StatusInternalServerError)
		return
	}

	err = database.SaveCotacao(cotacao)
	if err != nil {
		log.Printf("erro ao salvar cotacao: %v", err)
		http.Error(w, "nao foi possível salvar a cotacao no banco de dados", http.StatusInternalServerError)
		return
	}

	response := response.CotacaoResponse{
		Bid: cotacao,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func capturarCotacao() (string, error) {

	log.Println("Realizando a requisição para a API de cotação")

	ctx, cancel := context.WithTimeout(context.Background(), apiTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", urlCotacao, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("a aplicacao de cotacao retornou o status code: %v\n", resp.StatusCode)
		return "", err
	}

	var cotacao model.Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		return "", err
	}

	return cotacao.USDBRL.Bid, nil
}
