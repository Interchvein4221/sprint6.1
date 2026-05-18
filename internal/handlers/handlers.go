package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "../index.html")
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
	file, header, err := r.FormFile("file")
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
	rawTime := time.Now().UTC().String()
	rawTime = strings.ReplaceAll(rawTime, " ", "_")
	rawTime = strings.ReplaceAll(rawTime, ":", "-")
	ext := filepath.Ext(header.Filename)
	filename := rawTime + ext
	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "file error", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()
	_, _ = outFile.WriteString(result)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))

}
