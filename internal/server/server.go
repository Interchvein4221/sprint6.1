package server

import (
	"net/http"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func NewServer() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	return mux

}
