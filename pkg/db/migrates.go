package db

import (
	"github.com/celalettin-ceylan/go-bookstore/internal/models"
	"gorm.io/gorm"
)

func CreateBookTablesIfNotExist(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})
}

func CreateTables(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})
	// other models
}
