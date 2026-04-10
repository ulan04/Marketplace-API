package routes

import (
	"marketplace-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	
	r.GET("/items", handlers.GetItems)
	r.POST("/items", handlers.CreateItem)
	r.GET("/items/:id", handlers.GetItemByID)
	r.PUT("/items/:id", handlers.UpdateItem)
	r.DELETE("/items/:id", handlers.DeleteItem)

	
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUserByID)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)
}