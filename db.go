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

	instance, err := sql.Open("postgres", pgUrl)
	logFatal(err)

	db = instance

	err = db.Ping()
	logFatal(err)

	fmt.Println("database connection established succeefully!")
}

func CloseDBConn() {
	fmt.Println("Database connection closed!")
	db.Close()
}

func SelectFromDB(statement string) []ExistingPost {

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

	return results

}

func WriteToDB(statement string, args ...string) {
	stmt, err := db.Prepare(statement)
	logFatal(err)

	a := make([]interface{}, len(args))
	for i, v := range args {
		a[i] = v
	}

	defer stmt.Close()

	result, err := stmt.Exec(a...)
	logFatal(err)

	n, err := result.RowsAffected()

	fmt.Println(n, "row affected")
}
