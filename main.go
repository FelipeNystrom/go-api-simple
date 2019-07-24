package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	PostRoutes(s)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
