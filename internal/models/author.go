package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Book   []Book `json:"books" gorm:"foreignKey:AuthorID"`
}
