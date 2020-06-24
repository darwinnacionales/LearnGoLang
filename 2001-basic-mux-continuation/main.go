package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Item is a struct that groups all necessary fields
type Item struct {
	Data string `json:"data"`
}

var data []Item = []Item{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add", addItem).Methods("POST")

	err := http.ListenAndServe(":5000", router)
	log.Fatalln(err)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem)
	data = append(data, newItem)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
