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

func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	if err := r.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) GetBookById(id uint) (models.Book, error) {
	var book models.Book
	if err := r.DB.First(&book, id).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r *BookRepository) UpdateBook(book *models.Book) (*models.Book, error) {
	if err := r.DB.Save(book).Error; err != nil {
		return book, err
	}
	return book, nil
}
