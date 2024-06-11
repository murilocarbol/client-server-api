package server

import (
	"encoding/json"
	"github.com/murilocarbol/client-server-api/server/model"
	response2 "github.com/murilocarbol/client-server-api/server/response"
	"io"
	"log"
	"net/http"
)

const UrlCotacao = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

func SubirServer() {

	log.Println("Subindo servidor")

	router := http.NewServeMux()
	router.HandleFunc("/buscar_cotacao", handler)

	http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, r *http.Request) {

	log.Println("Realizando a requisição para a API de cotação")

	resp, err := http.Get(UrlCotacao)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response *model.Cotacao
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	var serverResponse = response2.ServerResponse{
		Bid: response.USDBRL.Bid,
	}

	respServer, err := json.Marshal(&serverResponse)
	if err != nil {
		panic(err)
	}

	w.Write(respServer)
}
