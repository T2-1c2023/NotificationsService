package main

import (
	dateGenerator "github.com/T2-1c2023/NotificationsService/app/utilities"
	_ "github.com/T2-1c2023/NotificationsService/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var creationDate string = dateGenerator.GetCurrentDate()

// GetStatus     godoc
// @Summary      Check the service's status.
// @Description  Returns a 200 code and JSON with a message.
// @Success      200
// @Router       / [get]
func getStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Notifications microservice running correctly",
	})
}

type HealthResponse struct {
	CreationDate string `json:"creation_date"`
	LastResponse string `json:"last_response"`
	Description  string `json:"description"`
}

// GetHealth     godoc
// @Summary      Get the service's health status.
// @Description  Returns a 200 code and JSON with the date the service started running and a description.
// @Success      200 {object} HealthResponse
// @Router       /health [get]
func getHealth(c *gin.Context) {
	response := HealthResponse{
		CreationDate: creationDate,
		LastResponse: dateGenerator.GetCurrentDate(),
		Description:  "Notifications microservice for FiuFit",
	}
	c.JSON(200, response)
}

func SetupRouter() *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	router.GET("/", getStatus)

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", getHealth)

	return router
}
