package handlers

import (
	"net/http"

	"github.com/blogster/types"
	"github.com/gin-gonic/gin"
)

// RegisterCommentRoutes defines author CRUD routes
func RegisterCommentRoutes(router *gin.RouterGroup) {
	router.POST("/:postID/react/", react())
	router.PATCH("/:postID/react/:reactID/", updateReact())
}

func react() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("postID")
		var payload types.ReactDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err = controllers.React(postID, payload.Reaction)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to comment"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Commented!"})
	}
}

func updateReact() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("postID")
		reactID := c.Param("reactID")
		var payload types.ReactDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err = controllers.UpdateReact(postID, reactID, payload.Reaction)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update comment"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Commented!"})
	}
}
