package client

import (
	"context"
	"encoding/json"
	"github.com/murilocarbol/client-server-api/client/model"
	"io"
	"net/http"
	"time"
)

func BuscarCotacao(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/buscar_cotacao", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response *model.ServerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(response.Bid))
}
