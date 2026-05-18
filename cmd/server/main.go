package main

import (
	"log"
	"net/http"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)
	func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	log.Println("server started on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}