package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	// Define a route handler for the root URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Notifications microservice running correctly",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Notifications microservice running correctly",
		})
	})

	return router
}
