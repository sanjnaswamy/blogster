package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/blogster/controllers"
	"github.com/blogster/types"
)

// RegisterAuthorRoutes defines author CRUD routes
func RegisterAuthorRoutes(router *gin.RouterGroup) {
	router.GET("/username/:username/available/", unameAvailability())
	router.POST("/signup/", registration())
	router.POST("/login/", login())
	router.PATCH("/:authorID/password/", updatePassword())

}

func unameAvailability() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Param("username")
		uname := strings.ToLower(userName)

		authorsList, err := controllers.GetAuthors()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get username list"})
			return
		}

		for _, value := range authorsList {
			if uname == value.UserName {
				c.JSON(http.StatusOK, gin.H{"available": false})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"available": true})
	}
}

func registration() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload types.RegisterAuthorDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err = controllers.Registration(payload.Name, payload.UserName, payload.Email, payload.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to register"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	}
}

func login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload types.LoginDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err = controllers.Login(payload.UserName, payload.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to login"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
	}
}

func updatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("authorID")
		var payload types.UpdatePasswordDTO
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		err = controllers.UpdatePassword(authorID, payload.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password updation successful"})
	}
}
