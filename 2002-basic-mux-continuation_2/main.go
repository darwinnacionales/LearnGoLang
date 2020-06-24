package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User is a struct
type User struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Post is a struct
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var data []Post = []Post{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add", addItem).Methods("POST")

	err := http.ListenAndServe(":5000", router)
	log.Fatalln(err)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	data = append(data, newPost)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
