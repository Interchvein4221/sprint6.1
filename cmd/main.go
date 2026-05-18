package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "morse: ", log.LstdFlags)
	srv := server.NewServer(logger)
	logger.Fatal(srv.ListenAndServe())

}
