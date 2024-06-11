package util

import (
	"net/http"
)

func SetupHealthCheck(router *http.ServeMux) {
	router.HandleFunc("/health", healthCheck)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is up to date"))
}
