package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//postgres://bjrjiwkb:YFixAZUxZzWfXF7qClfXVrs-bUj297My@ziggy.db.elephantsql.com:5432/bjrjiwkb

var db *sql.DB
var err error

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	user     = "bjrjiwkb"
	password = "YFixAZUxZzWfXF7qClfXVrs-bUj297My"
	host     = "ziggy.db.elephantsql.com"
	port     = 5432
	dbname   = "bjrjiwkb"
	sslmode  = "verify-full"
)

func main() {

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", user, password, host, port, dbname, sslmode)

	//connStr := "user=1234 password=gfgdfg-bUj297My host=11ziggy.db.elephantsql.com port=5432 dbname=bjrjiwkb sslmode=verify-full"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//logFatal(err)

	db.Ping()
	fmt.Println("DB connection successful===")

	router := mux.NewRouter()

	router.HandleFunc("/getbooks", getBooks).Methods("GET")
	router.HandleFunc("/postbooks", addBooks).Methods("POST")
	router.HandleFunc("/pathcbooks", updateBooks).Methods("PATCH")
	router.HandleFunc("/deletebooks", removeBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//Book is
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	books := []Book{}

	rows, err := db.Query("select * from library")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Name, &book.Author)
		logFatal(err)
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
	fmt.Println("ruban getbooks func")
}

//gorilla mux, gin, fasthttp, treemux

//addBook, updateBook, removeBook

func addBooks(w http.ResponseWriter, r *http.Request) {
	db.Query("")
	fmt.Println("Added a new Book")
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updated an existing Book")
}

func removeBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("DELETE FROM library WHERE book_id = 34")
	logFatal(err)
	defer rows.Close()
	for rows.Next() {
	}
	fmt.Println("Deleted the Book records")
}
