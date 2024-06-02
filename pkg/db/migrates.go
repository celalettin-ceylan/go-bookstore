package db

import (
	"github.com/celalettin-ceylan/go-bookstore/internal/models"
	"gorm.io/gorm"
)

func CreateTables(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})
}
