package controllers

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/blogster/database"
	"github.com/blogster/models"
)

// GetAuthors gets all user names
func GetAuthors() (*models.Author, error) {
	authors, err := database.GetAuthors()
	if err != nil {
		return nil, errors.New("Failed to get list of authors")
	}

	return &authors, nil
}

// Registration signs up a user
func Registration(name, userName, email, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Unable to hash password using bcrypt")
	}

	err := database.Register(&models.Author{
		Name:     name,
		UserName: userName,
		Email:    email,
		Password: string(password),
	})

	if err != nil {
		log.Println{err}
		return err
	}

	return nil
}

// Login checks if user is valid and logs in
func Login(userName, password string) error {
	author, err := database.GetAuthor(&models.Author{UserName: username})
	if err != nil {
		return errors.New("Failed to get list of authors")
	}

	bytePassword := []byte(password)
	byteHashedPassword := []byte(author.Password)
	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return errors.New("Invalid password")
	}

	return nil
}

// UpdatePassword updates the account password
func UpdatePassword(authorID, newPassword string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Unable to hash password using bcrypt")
	}

	err := database.UpdateAuthor(&models.Author{Password: passwordHash})
	if err != nil {
		return errors.New("Failed to update password")
	}

	return nil
}
