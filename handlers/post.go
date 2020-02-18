package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blogster/controllers"
	"github.com/blogster/types"
)

// RegisterPostRoutes defines post routes
func RegisterPostRoutes(router *gin.RouterGroup) {
	router.GET("/", getAllPosts())
	router.GET("/:postID/", getPost())
}

// RegisterAuthorPostRoutes defines post CRUD routes
func RegisterAuthorPostRoutes(router *gin.RouterGroup) {
	router.POST("/", post())
	router.GET("/", getAllPosts())
	router.GET("/", getPostsListForAuthor())
	router.PATCH("/:postID/", updatePost())
	router.DELETE("/:postID/", deletePost())
}

func getAllPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		postsList, err = controllers.GetAllPosts()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get posts list"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"postsList": postsList})
	}
}

func post() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("authorID")
		var payload types.PostDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err = controllers.Post(authorID, payload.Content)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Posted!"})
	}
}

func getPostsListForAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("authorID")
		postsList, err = controllers.GetPostsList(authorID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get posts list"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"postsList": postsList})
	}
}

func getPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("postID")
		var payload types.PostDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		post, err = controllers.GetPost(postID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"post": post})
	}
}

func updatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("authorID")
		postID := c.Param("postID")
		var payload types.PostDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err = controllers.UpdatePost(authorID, postID, payload.Name, payload.Content)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
	}
}

func deletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("authorID")
		postID := c.Param("postID")

		err = controllers.UpdatePost(authorID, postID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post deleted!"})
	}
}
