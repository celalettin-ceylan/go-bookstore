package repositories

import (
	"fmt"
	"log"
	"testing"

	"github.com/celalettin-ceylan/go-bookstore/internal/models"
	"github.com/celalettin-ceylan/go-bookstore/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=cc password=12345 dbname=gobookstore_test port=5432 sslmode=disable"

func connectDatabase() *gorm.DB {
	var DB *gorm.DB
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	db.CreateTables(DB)
	return DB
}

func TestCreateBook(t *testing.T) {
	book := models.Book{
		Model:    gorm.Model{},
		AuthorID: 2,
		Author:   models.Author{Gender: "Female", Name: "Merve CEYLAN"},
		Title:    "Bati Expressinde Cinayet2",
		Price:    50,
	}

	DB := connectDatabase()
	NewBookRepository(DB).CreateBook(&book)
}

func TestGetAllBook(t *testing.T) {
	DB := connectDatabase()
	books, err := NewBookRepository(DB).GetAllBooks()
	if err != nil {
		fmt.Printf("Could not get data")
		return
	}
	fmt.Print(books)
}

func TestGetBookById(t *testing.T) {
	DB := connectDatabase()
	book, err := NewBookRepository(DB).GetBookById(1)
	if err != nil {
		fmt.Printf("Could not get data")
		return
	}
	fmt.Print(book)
}

func TestUpdateBook(t *testing.T) {
	book := models.Book{
		Model:    gorm.Model{ID: 1},
		Title:    "Kuzey Expressinde Cinayet",
		AuthorID: 1,
		Author:   models.Author{Model: gorm.Model{ID: 1}, Gender: "Male", Name: "Celalettin ISLAM"},
		Price:    40,
	}

	DB := connectDatabase()
	updatedBook, err := NewBookRepository(DB).UpdateBook(&book)
	if err != nil {
		fmt.Print("Can not update book record")
	}
	fmt.Print(updatedBook)
}

func TestDeleteBook(t *testing.T) {
	DB := connectDatabase()
	err := NewBookRepository(DB).DeleteBook(1)
	if err != nil {
		fmt.Printf("Can not delete book")
	}
	fmt.Printf("Deleted successfully book")
}
