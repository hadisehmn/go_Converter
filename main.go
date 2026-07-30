package main

import (
	"log"
	"net/http"

	"go-practice/CONVERTER/controller"
	service "go-practice/CONVERTER/services"
)

func main() {

	convertService := service.NewConvertService()
	convertController := controller.NewConvertController(convertService)
	http.HandleFunc("/convert", convertController.ConvertFile)

	log.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
