package main

import (
	"log"

	"github.com/celalettin-ceylan/go-bookstore/pkg/db"
)

func main() {
	log.Println("Hello World!")
	postgres_db := db.Setup()
	db.CreateTables(postgres_db)
}
