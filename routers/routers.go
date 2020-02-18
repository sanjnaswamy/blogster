package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blogster/handlers"
)

// RegisterRoutes returns a registered gin engine
func RegisterRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/healthz/", func(c *gin.Context) {
		c.String(200, "healthy")
	})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api")
	v1 := api.Group("v1")

	author := v1.Group("author")
	handlers.RegisterAuthorRoutes(author)

	post := v1.Group("post")
	handlers.RegisterPostRoutes(post)

	authorPost := v1.Group("/:authorID/post")
	handlers.RegisterPostRoutes(authorPost)

	react := post.Group("react")
	handlers.RegisterReactRoutes(react)

	comment := authorPost.Group("comment")
	handlers.RegisterCommentRoutes(comment)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "PAGE_NOT_FOUND"})
	})

	return router
}
