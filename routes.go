package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func handleBadRequest(err error, w http.ResponseWriter) {
	if err != nil {
		// Handle error
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

var decoder = schema.NewDecoder()

func PostRoutes(r *mux.Router) {

	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/post", createPost).Methods("POST")
	r.HandleFunc("/post/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/post/{id}", deletePost).Methods("DELETE")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := SelectFromDB("select * from posts;")

	posts, _ := json.Marshal(data)

	w.Write(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post PostBody

	err := r.ParseForm()

	handleBadRequest(err, w)

	err = decoder.Decode(&post, r.Form)

	handleBadRequest(err, w)

	WriteToDB("INSERT INTO posts (title, body, author) VALUES($1, $2,$3);", post.Title, post.Body, post.Author)

	w.WriteHeader(http.StatusOK)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idMap := mux.Vars(r)
	var post PostBody

	err := r.ParseForm()

	handleBadRequest(err, w)

	err = decoder.Decode(&post, r.Form)

	handleBadRequest(err, w)

	WriteToDB("UPDATE posts SET title = $2, body = $3, author = $4 WHERE id = $1;", idMap["id"], post.Title, post.Body, post.Author)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idMap := mux.Vars(r)
	WriteToDB("DELETE FROM posts WHERE id = $1 RETURNING *", idMap["id"])

	w.WriteHeader(http.StatusOK)
}
