package controllers

import (
	"github.com/blogster/database"
	"github.com/blogster/models"
)

// GetAllPosts defines fetching all existing posts by all authors
func GetAllPosts() ([]models.Post, error) {
	posts, err := database.GetPosts(&models.Post)
	return posts, err
}

// Post defines blog post creation
func Post(authorID int, content string) error {
	err := database.Post(&models.Post{
		AuthorID: authorID,
		Content:  content,
	})
	return err
}

// GetPostsListForAuthor defines getting list of post of that author
func GetPostsListForAuthor(authorID int) (posts, error) {}
