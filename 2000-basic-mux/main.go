package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var data []string = []string{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test1", test1)
	router.HandleFunc("/test2", test2)
	router.HandleFunc("/add/{item}", addItem).Methods("POST")

	err := http.ListenAndServe(":5000", router)
	log.Fatalln(err)
}

func test1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a test"))
}

func test2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Name string
		Age  int
	}{
		Name: "Darwin",
		Age:  28,
	})
}

func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	routeVariable := mux.Vars(r)["item"]

	data = append(data, routeVariable)

	json.NewEncoder(w).Encode(data)
}
