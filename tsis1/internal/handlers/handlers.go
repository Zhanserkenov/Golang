package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ваш_логин/Golang/internal/models"
)

var friends = []models.Person{
	{Name: "John", Age: 30},
	{Name: "Alice", Age: 25},
	// Добавьте больше друзей по мере необходимости
}

func GetFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

func GetFriendByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	name := params["name"]

	for _, friend := range friends {
		if friend.Name == name {
			json.NewEncoder(w).Encode(friend)
			return
		}
	}

	http.NotFound(w, r)
}
