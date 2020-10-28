package main

import (
	"log"
	"net/http"
	"github.com/KiranGosavi/gopractice/book"
)

const basePath = "/api"

func main() {

	//basePath :=
	//book.SetupRoutes()

	book.SetupRoutes(basePath)
	err :=http.ListenAndServe(":5005", nil)
	if err != nil {
		log.Fatal(err)
	}
}