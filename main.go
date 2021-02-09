package main

import (
	"log"
	"net/http"

	"github.com/Garius6/blog/internal/server"
	"github.com/Garius6/blog/internal/storage/sqlstorage"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	store, _ := sqlstorage.New("articles.db")
	server := server.New(store)
	r := server.ConfigureRoutes()
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
