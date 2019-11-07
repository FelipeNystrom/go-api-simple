package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	Init()
	defer CloseDBConn()

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	PostRoutes(s)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
