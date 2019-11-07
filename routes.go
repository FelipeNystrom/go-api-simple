package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var posts []ExistingPost

func PostRoutes(r *mux.Router) {

	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/post", createPost).Methods("POST")
	r.HandleFunc("/post/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/post/{id}", deletePost).Methods("DELETE")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts := FetchFromDB("select * from posts;")
	fmt.Println(posts)
	// w.Write(posts)
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

	var post ExistingPost

	err = decoder.Decode(&post, r.Form)

	if err != nil {
		// Handle error
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	posts = append(posts, post)

	w.WriteHeader(http.StatusOK)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()

	params := mux.Vars(r)

	Id, _ := strconv.Atoi(params["id"])

	if err != nil {
		// Handle error
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	decoder := schema.NewDecoder()

	for index, item := range posts {
		if item.Id == Id {
			posts = append(posts[:index], posts[index+1:]...)

			var post ExistingPost
			_ = decoder.Decode(&post, r.Form)

			post.Id = Id

			posts = append(posts, post)
			json.NewEncoder(w).Encode(&post)
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
