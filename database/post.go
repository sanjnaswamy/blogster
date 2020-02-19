package database

import (
	"errors"

	"github.com/blogster/models"
)

// GetPosts gets list of all posts
func GetPosts(condition interface{}) ([]models.Post, error) {
	database := GetDb()
	var result []models.Post

	err := database.Where(condition).Find(&result).Error
	if err != nil {
		return nil, errors.New("No posts found")
	}

	return result, nil
}

// Post defines blog post creation
func Post(data interface{}) error {
	database := GetDb()
	err := database.Create(data).Error
	return err
}

// GetPost gets the post object by condition
func GetPost(condition interface{}) (*models.Post, error) {
	database := GetDb()
	var result models.Post

	err := database.Where(condition).Find(&result).Error
	if err != nil {
		return nil, errors.New("No post found")
	}

	return result, nil
}

// UpdatePost defines post updation
func UpdatePost(condition interface{}, fields interface{}) error {
	database := GetDb()
	var post models.Post
	if database.Where(condition).Find(&post).RecordNotFound() {
		return errors.New
	}

	err := database.Model(&post).Where(condition).Updates(fields).Error
	if err != nil {
		return errors.New("Failed to update post")
	}
	return nil
}

// DeletePost defines post deletion
func DeletePost(condition interface{}) error {
	database := GetDb()
	err := database.Unscoped().Delete(condition).Error
	if err != nil {
		return erros.New("Failed to delete post")
	}

	return nil
}
