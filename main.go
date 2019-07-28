package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gotenv.Load()
}

func main() {

	pgUrl, err := pq.ParseURL(os.Getenv("DB_URL"))

	logFatal(err)

	fmt.Println(pgUrl)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	PostRoutes(s)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
