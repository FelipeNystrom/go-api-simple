package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

var posts []Post

func PostRoutes(router *mux.Router) {

	posts = append(posts, Post{Id: 0, Title: "Hello world", Author: "Jane Doe", Text: "This is the first post from go test api!"}, Post{Id: 1, Title: "Hello world", Author: "Jane Doe", Text: "This is the first post from go test api!"})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/post", createPost).Methods("POST")
	router.HandleFunc("/post/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()

	if err != nil {
		// Handle error
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	decoder := schema.NewDecoder()

	var post Post

	err = decoder.Decode(&post, r.Form)

	if err != nil {
		// Handle error
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	posts = append(posts, post)
	fmt.Println(posts)

	w.WriteHeader(http.StatusOK)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	Id, _ := strconv.Atoi(params["id"])

	for index, item := range posts {
		if item.Id == Id {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.Id = Id
			posts = append(posts, post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)

}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Id, err := strconv.Atoi(params["id"])

	fmt.Println(Id, err)
	for i, item := range posts {

		if item.Id == Id {
			fmt.Println(i, item)
			posts = append(posts[:i], posts[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}
