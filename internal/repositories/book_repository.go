package repositories

import (
	"log"

	"github.com/celalettin-ceylan/go-bookstore/internal/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) CreateBook(book *models.Book) (*models.Book, error) {
	if err := r.DB.Create(book).Error; err != nil {
		log.Printf("Error Creating Book %v", err)
		return book, err
	}
	return book, nil
}
