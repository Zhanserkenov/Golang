package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ваш_логин/mywebapp/internal/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/friends", handlers.GetFriends).Methods("GET")
	router.HandleFunc("/friends/{name}", handlers.GetFriendByName).Methods("GET")
	router.HandleFunc("/health", HealthCheck).Methods("GET")

	return router
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Мое веб-приложение работает нормально!\nАвтор: Ваше Имя"))
}
