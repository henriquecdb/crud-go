package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Read(w http.ResponseWriter, r *http.Request) {
	
}

func Create(w http.ResponseWriter, r *http.Request) {
	
}

var db *sql.DB

func init() {
	var err error
	db, err := sql.Open("postgres", "postgres://henrique:123456@postgres/cruds?sslmode=disable")
	if err != nil {
		panic(err)
	}
	
	if err = db.Ping(); err != nil {
		panic(err)
	}
	
	fmt.Println("Connection established.")
}

func main() {
	http.HandleFunc("/users/read", Read)
	http.HandleFunc("/users/create", Create)
	http.ListenAndServe(":8080", nil)
}
