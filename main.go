package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	Id    int
	Name  string
	Email string
	Age   int
}

func Read(w http.ResponseWriter, r *http.Request) {

}

func Create(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    u := User{}
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        http.Error(w, "Server failed to handle the request.", http.StatusInternalServerError)
        return
    }

    _, err = db.Exec("INSERT INTO users (name, email, age) VALUES ($1,$2,$3)", u.Name, u.Email, u.Age)
    if err != nil {
        http.Error(w, "Server failed to handle the request.", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "postgres://henrique:123456@postgres/crud?sslmode=disable")
    if err != nil {
        panic(err)
    }

    if err = db.Ping(); err != nil {
        panic(err)
    }

    fmt.Println("Connection established.")
}

func main() {
	// http.HandleFunc("/users/read", Read)
	http.HandleFunc("/users/create", Create)
	http.ListenAndServe(":8081", nil)
}
