package app

import (
	"anggi.blog/controllers/posts"
	"anggi.blog/controllers/users"
	"github.com/gin-gonic/gin"
)

func MapUrls() {
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})

	apiRouter := router.Group("/api")
	{
		apiRouter.POST("/register", users.Register)
		apiRouter.POST("/login", users.Login)
		apiRouter.GET("/user", users.GetUser)
		apiRouter.GET("/logout", users.Logout)

		blogRouter := apiRouter.Group("/blog")
		{
			blogRouter.GET("/", posts.GetPosts)
			blogRouter.POST("/", posts.CreatePost)
			blogRouter.GET("/:id", posts.GetPost)
			blogRouter.PUT("/:id", posts.UpdatePost)
			blogRouter.DELETE("/:id", posts.DeletePost)
		}
	}
}
