package repositories

import (
	"fmt"
	"testing"

	"github.com/celalettin-ceylan/go-bookstore/internal/models"
	"gorm.io/gorm"
)

func TestCreateAuthor(t *testing.T) {
	author := models.Author{
		Model:  gorm.Model{},
		Gender: "Male",
		Name:   "Umit Karakaya",
	}

	DB := connectDatabase()
	NewAuthorRepository(DB).CreateAuthor(&author)
}

func TestGetAuthorById(t *testing.T) {
	db := connectDatabase()
	author, err := NewAuthorRepository(db).GetAuthorById(3)

	if err != nil {
		fmt.Printf("Error during get auth: %v", err)
		return
	}
	fmt.Println(author)
}

func TestGetAllAuthor(t *testing.T) {
	db := connectDatabase()
	authors, err := NewAuthorRepository(db).GetAllAuthor()
	if err != nil {
		fmt.Printf("Error during get all author %v", err)
		return
	}
	fmt.Println(authors)
}

func TestDeleteAuthor(t *testing.T) {
	db := connectDatabase()
	if err := NewAuthorRepository(db).DeleteAuthor(1); err != nil {
		fmt.Printf("Can not delete author %v", err)
		return
	}
	fmt.Printf("Deleted successfully")
}

func TestUpdateAuthor(t *testing.T) {

	DB := connectDatabase()
	ar := NewAuthorRepository(DB)

	updatedAuthor := models.Author{
		Model:  gorm.Model{ID: 1},
		Gender: "Female",
		Name:   "Updated CM",
	}

	responseAuthor, err := ar.UpdateAuthor(&updatedAuthor)

	if err != nil {
		fmt.Printf("TestUpdateAuthor Error : %v", err)
		return
	}

	fmt.Println(responseAuthor)
}
