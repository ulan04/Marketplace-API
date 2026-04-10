package main

import (
	"log"
	"marketplace-api/config"
	"marketplace-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.RegisterRoutes(r)

	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}