package main

import (
	"log"
	"net/http"

	"github.com/celalettin-ceylan/go-bookstore/internal/bookstore/routes"
	"github.com/celalettin-ceylan/go-bookstore/pkg/db"
)

func main() {

	DB := db.Setup()
	db.CreateBookTablesIfNotExist(DB)

	router := routes.SetupBookRoutes()

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
