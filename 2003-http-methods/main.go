package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", addItem).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)
	log.Fatalln(err)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	postID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid ID provided as parameter"))
		return
	}

	for idx, p := range posts {
		if p.ID == postID {
			posts = append(posts[:idx], posts[idx+1:]...)
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("Post with given ID not found"))
}

func patchPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	postID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid ID provided as parameter"))
		return
	}

	for idx, p := range posts {
		if p.ID == postID {
			json.NewDecoder(r.Body).Decode(&posts[idx])
			json.NewEncoder(w).Encode(posts[idx])
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("Post with given ID not found"))
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postID, err := strconv.Atoi(mux.Vars(r)["id"])
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid ID provided as parameter"))
		return
	}

	for idx, p := range posts {
		if p.ID == postID {
			posts[idx] = updatedPost

			json.NewEncoder(w).Encode(posts[idx])
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("Post with given ID not found"))
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	postID, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid ID provided as parameter"))
		return
	}

	for _, p := range posts {
		if p.ID == postID {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("Post with given ID not found"))
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
