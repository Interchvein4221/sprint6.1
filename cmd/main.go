package main

import (
	"log"
	"net/http"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)
	

func main() {

	srv := server.NewServer()
	log.Println("server started :8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}

}
