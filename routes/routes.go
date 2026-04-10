package routes

import (
	"marketplace-api/handlers"
	"marketplace-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.GET("/items", handlers.GetItems)
	r.GET("/items/:id", handlers.GetItemByID)
	r.GET("/categories", handlers.GetCategories)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
	
		protected.POST("/items", handlers.CreateItem)
		protected.PUT("/items/:id", handlers.UpdateItem)
		protected.DELETE("/items/:id", handlers.DeleteItem)

		protected.GET("/users", handlers.GetUsers)
		protected.GET("/users/:id", handlers.GetUserByID)

		protected.POST("/categories", handlers.CreateCategory)
	}
}