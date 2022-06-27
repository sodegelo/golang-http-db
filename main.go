package main

import (
	"database/sql"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name        string
	Description string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	product := Product{Name: "Julio", Description: "Julio"}
	persistProduct(product)
	w.Write([]byte("Go ON.."))
}

func persistProduct(product Product) {
	db, err := sql.Open("sqlite3", "production.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO products(name, description) values($1,$2)")

	if err != nil {
		panic(err)
	}
	stmt.Exec(product.Name, product.Description)
}
