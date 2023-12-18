package rest

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func StartServer(addr string) {
	http.HandleFunc("/", handler)
	http.Handle("/style.css", http.FileServer(http.Dir("/app/index")))
	http.ListenAndServe(addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	filePath, err := filepath.Abs("index/index.html")
	if err != nil {
		log.Printf("Error getting absolute file path: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Чтение содержимого файла
	stat, err := file.Stat()
	if err != nil {
		log.Printf("Error getting file info: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}
