package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func PostRoutes(r *mux.Router) {

	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/post", createPost).Methods("POST")
	r.HandleFunc("/post", updatePost).Methods("PUT")
	r.HandleFunc("/post", deletePost).Methods("DELETE")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := SelectFromDB("select * from posts;")

	posts, _ := json.Marshal(data)

	w.Write(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()

	if err != nil {
		// Handle error
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var post PostBody

	err = decoder.Decode(&post, r.Form)

	if err != nil {
		// Handle error
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	WriteToDB("INSERT INTO posts (title, body, author) VALUES($1, $2,$3);", post.Title, post.Body, post.Author)

	fmt.Println(r.Form)

	w.WriteHeader(http.StatusOK)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// err := r.ParseForm()

}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// err := r.ParseForm()
}
