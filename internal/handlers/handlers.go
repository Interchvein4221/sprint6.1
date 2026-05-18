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

func UploadHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {

        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

        return

    }

    if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {

        http.Error(w, "invalid content type", http.StatusBadRequest)

        return

    }

    err := r.ParseMultipartForm(10 << 20)

    if err != nil {

        http.Error(w, "invalid form", http.StatusBadRequest)

        return

    }

    file, header, err := r.FormFile("file")

    if err != nil {

        http.Error(w, err.Error(), http.StatusBadRequest)

        return

    }

    defer file.Close()

    data, err := io.ReadAll(file)

    if err != nil {

        http.Error(w, err.Error(), http.StatusInternalServerError)

        return

    }

    input := strings.TrimSpace(string(data))

    result, err := service.Convert(input)

    if err != nil {

        http.Error(w, err.Error(), http.StatusBadRequest)

        return

    }

    os.MkdirAll("./uploads", 0755)

    ext := filepath.Ext(header.Filename)

    if ext == "" {

        ext = ".txt"

    }

    filename := "./uploads/" + time.Now().UTC().Format("20060102150405") + ext

    err = os.WriteFile(filename, []byte(result), 0644)

    if err != nil {

        http.Error(w, err.Error(), http.StatusInternalServerError)

        return

    }

    w.Header().Set("Content-Type", "text/plain; charset=utf-8")

    w.Write([]byte(result))

}