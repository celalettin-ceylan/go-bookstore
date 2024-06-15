package repositories

import (
	"github.com/celalettin-ceylan/go-bookstore/internal/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{DB: db}
}

func (ar *AuthorRepository) CreateAuthor(author *models.Author) (*models.Author, error) {

	if err := ar.DB.Create(&author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func (ar *AuthorRepository) GetAllAuthor() (*[]models.Author, error) {
	var authors []models.Author

	if err := ar.DB.Find(&authors).Error; err != nil {
		return nil, err
	}
	return &authors, nil
}

func (ar *AuthorRepository) GetAuthorById(id uint) (*models.Author, error) {
	var author models.Author
	if err := ar.DB.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

// Created_At column can be changed after update or delete operations because of using gorm.Model attributes.
// So we should get old data or model before update or delete record.
func (ar *AuthorRepository) UpdateAuthor(author *models.Author) (*models.Author, error) {

	oldAuthor, err := ar.GetAuthorById(author.ID)

	if err != nil {
		return nil, err
	}

	author.Model.CreatedAt = oldAuthor.CreatedAt

	if err := ar.DB.Save(author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

// Created_At column can be changed after update or delete operations because of using gorm.Model attributes.
// So we should get old data or model before update or delete record.
func (ar *AuthorRepository) DeleteAuthor(id uint) error {

	oldAuthor, err := ar.GetAuthorById(id)

	if err != nil {
		return err
	}

	if err := ar.DB.Delete(&oldAuthor, id).Error; err != nil {
		return err
	}
	return nil
}
