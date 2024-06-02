package models

type Book struct {
	ID       uint    `gorm:"primaryKey"`
	Title    string  `gorm:"size:255;not null"`
	Author   string  `gorm:"size:255;not null"`
	Price    float64 `gorm:"not null"`
	Quantity int     `gorm:"not null"`
}
