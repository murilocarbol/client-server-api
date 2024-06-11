package main

import (
	"github.com/murilocarbol/client-server-api/util"
	"log"
	"net/http"

	"github.com/murilocarbol/client-server-api/client"
	"github.com/murilocarbol/client-server-api/server"
)

func main() {

	// sobe o server para capturar cotacao
	go server.SubirServer()

	routerMux := http.NewServeMux()
	routerMux.HandleFunc("/cotacao", handleRequest)

	go util.SetupHealthCheck(routerMux)

	http.ListenAndServe(":3000", routerMux)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for client:")
	client.BuscarCotacao(w, r)
}
