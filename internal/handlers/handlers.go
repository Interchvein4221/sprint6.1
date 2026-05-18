package handlers

import (
	"io"
	"net/http"
     "time"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"strings"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "index.html")

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "bad form", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "no file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "read error", http.StatusInternalServerError)
		return
	}
	result := service.Convert(string(data))
	name := strings.ReplaceAll(time.Now().UTC().Format(time.RFC3339Nano), ":", "-")
	filename := name + ".txt"
	_ = os.WriteFile(filename, []byte(result), 0644)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(result))

}