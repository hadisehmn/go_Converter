package main

import (
	"log"
	"net/http"

	converter "go-practice/CONVERTER/internal/converter"
	"go-practice/CONVERTER/internal/user"
)

func main() {

	convertService := converter.NewConvertService()
	convertController := user.NewConvertController(convertService)
	http.HandleFunc("/convert", convertController.ConvertFile)

	http.Handle(
		"/downloads/", http.StripPrefix(
			"/downloads/", http.FileServer(http.Dir("./downloads")),
		),
	)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is healthy"))
	})

	converter.StartCleanupJob()
	log.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
