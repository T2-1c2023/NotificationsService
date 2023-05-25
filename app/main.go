// @title		Notifications Microservice API
// @version         1.0
// @description     Notifications microservice for FiuFit's backend.

package main

import (
	"os"

	routes "github.com/T2-1c2023/NotificationsService/app/routes"
)

func main() {
	router := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002" // Default port
	}

	router.Run(":" + port)
}
