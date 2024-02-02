package main

import (
	"log"
	"net/http"

	"github.com/ваш_логин/mywebapp/internal/server"
)

func main() {
	router := server.NewRouter()

	log.Println("Запуск сервера на :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
