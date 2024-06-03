package repositories

import (
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

func TestCreateBook(t *testing.T) {
	var DB *gorm.DB
	var err error

	book := models.Book{
		Title:    "Kuzey Expressinde Cinayet",
		Author:   "Agathe Cristhine",
		Price:    50,
		Quantity: 100,
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	db.CreateBookTablesIfNotExist(DB)
	NewBookRepository(DB).CreateBook(&book)
}
