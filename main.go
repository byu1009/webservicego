package main

import (
	"log"
	"webservicego/config"
	"webservicego/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal load file .env")
	}
	
	config.LoadConfig()
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	routes.SetupRoutes(r)

	// Custom 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    404,
			"message": "Route not found",
		})
	})

	r.Run(":8080")
}
