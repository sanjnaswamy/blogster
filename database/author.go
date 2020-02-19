package database

import (
	"errors"

	"github.com/blogster/models"
)

// GetAuthors gets list of all authors
func GetAuthors(condition interface{}) ([]models.Author, error) {
	database := GetDb()
	var result []models.Author

	err := database.Where(condition).Find(&result).Error
	if err != nil {
		return nil, errors.New("No authors found")
	}

	return result, nil
}

// GetAuthor gets list of all authors
func GetAuthor(condition interface{}) (*models.Author, error) {
	database := GetDb()
	var result models.Author

	err := database.Where(condition).Find(&result).Error
	if err != nil {
		return nil, errors.New("No author found")
	}

	return result, nil
}

// Register creates one instance of a given data
func Register(data interface{}) error {
	database := GetDb()
	err := database.Create(data).Error
	return err
}
