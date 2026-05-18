package server
<<<<<<< HEAD
=======
import (

	"log"
	"net/http"
	"time"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"

)

func NewServer(logger *log.Logger) *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	return &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

}
>>>>>>> bb32e31 (one18.05.2026)
