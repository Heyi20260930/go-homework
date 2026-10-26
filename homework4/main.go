package main


import (
	"homework4/config"
	"homework4/database"
	"homework4/handlers"
	"homework4/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(middleware.Logger())

	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		posts := api.Group("/posts")
		posts.GET("", handlers.ListPosts)
		posts.GET("/:id", handlers.GetPost)
		posts.GET("/:id/comments", handlers.ListComments)

		auth := posts.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.POST("", handlers.CreatePost)
			auth.PUT("/:id", handlers.UpdatePost)
			auth.DELETE("/:id", handlers.DeletePost)
			auth.POST("/:id/comments", handlers.CreateComment)
		}
	}

	if err := r.Run(config.ServerPort); err != nil {
		panic(err)
	}
}