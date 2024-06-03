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

// type Book struct {
// 	ID       uint    `gorm:"primaryKey"`
// 	Title    string  `gorm:"size:255;not null"`
// 	Author   string  `gorm:"size:255;not null"`
// 	Price    float64 `gorm:"not null"`
// 	Quantity int     `gorm:"not null"`
// }

var dsn = "host=localhost user=cc password=12345 dbname=gobookstore_test port=5432 sslmode=disable"

func connectDatabase() *gorm.DB {
	var DB *gorm.DB
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	db.CreateBookTablesIfNotExist(DB)
	return DB
}

func TestCreateBook(t *testing.T) {
	book := models.Book{
		Title:    "Kuzey Expressinde Cinayet",
		Author:   "Agathe Cristhine",
		Price:    50,
		Quantity: 100,
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
		ID:       2,
		Title:    "Kuzey Expressinde Cinayet",
		Author:   "Agathe Cristhine",
		Price:    35,
		Quantity: 100,
	}

	DB := connectDatabase()
	updatedBook, err := NewBookRepository(DB).UpdateBook(&book)
	if err != nil {
		fmt.Print("Can not update book record")
	}
	fmt.Print(updatedBook)
}
