package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"tsis1/models"
)

var creatures = []models.Creature{
	{Name: "Hippogriff", Description: "A majestic creature with the front legs, wings, and head of a giant eagle, and the body, hind legs, and tail of a horse."},
	{Name: "Thestral", Description: "A black, skeletal, bat-winged horse, visible only to those who have witnessed death."},
}

func GetCreatureList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creatures)
}

func GetCreature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["name"]
	for _, creature := range creatures {
		if creature.Name == name {
			json.NewEncoder(w).Encode(creature)
			return
		}
	}
	http.Error(w, "Creature not found", http.StatusNotFound)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a Harry Potter-themed web app. Created by [Arsen Zhanserkenov]")
}
