package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init() {

	pgUrl, err := pq.ParseURL(os.Getenv("DB_URL"))
	dbConn, err := sql.Open("postgres", pgUrl)
	logFatal(err)

	db = dbConn

	err = db.Ping()
	logFatal(err)

	fmt.Println("database connection established succeefully!")
}

func FetchFromDB(statement string) []ExistingPost {

	rows, err := db.Query(statement)
	if err != nil {
		panic(err)
	}

	var results []ExistingPost

	for rows.Next() {
		var id int
		var title string
		var body string
		var author string

		err := rows.Scan(&id, &title, &body, &author)

		logFatal(err)

		results = append(results, ExistingPost{id, PostBody{title, body, author}})
	}

	fmt.Println(results)

	return results

}

func Query(statement string, args []string) *sql.Rows {

	rows, err := db.Query(statement, args)
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)

	return rows

}

func CloseDBConn() {
	fmt.Println("Database connection closed!")
	db.Close()
}
