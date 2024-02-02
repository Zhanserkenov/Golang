package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tsis1/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/creatures", handlers.GetCreatureList).Methods("GET")
	r.HandleFunc("/creatures/{name}", handlers.GetCreature).Methods("GET")
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
