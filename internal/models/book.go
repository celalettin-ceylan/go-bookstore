package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string  `json:"title" gorm:"size:255;not null"`
	AuthorID uint    `json:"author_id" gorm:"not null"`
	Author   Author  `json:"author"`
	Price    float64 `json:"price" gorm:"not null"`
}
