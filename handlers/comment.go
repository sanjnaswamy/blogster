package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterCommentRoutes defines author CRUD routes
func RegisterCommentRoutes(router *gin.RouterGroup) {
	router.POST("/:postID/comment/", comment())
	router.PATCH("/:postID/comment/:commentID/", updateComment())
}

func comment() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("authorID")
		postID := c.Param("postID")
		var payload types.CommentDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}
		err = controllers.Comment(authorID, postID, payload.Comment)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to comment"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Commented!"})
	}
}

func updateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("authorID")
		postID := c.Param("postID")
		commentID := c.Param("commentID")
		var payload types.CommentDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err = controllers.UpdateComment(authorID, postID, commentID, payload.Comment)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update comment"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Commented!"})
	}
}
