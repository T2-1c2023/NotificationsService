// @title		Notifications Microservice API
// @version         1.0
// @description     Notifications microservice for FiuFit's backend.

package main

import (
	"os"
)

func main() {

	// Create a new Gin router
	router := SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002" // Default port
	}

	// Run the server
	router.Run(":" + port)
}
