package main

import (
	"github.com/murilocarbol/client-server-api/client"
	"github.com/murilocarbol/client-server-api/server"
)

func main() {

	// Sobe o server para capturar cotacao
	go server.SubirServer()

	// Chama o client para buscar cotacao
	client.BuscarCotacao()
}
