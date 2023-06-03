// @title		Notifications Microservice API
// @version         1.0
// @description     Notifications microservice for FiuFit's backend.

package main

import (
	"os"

	"github.com/T2-1c2023/NotificationsService/app/controller"
	routes "github.com/T2-1c2023/NotificationsService/app/routes"
	"github.com/T2-1c2023/NotificationsService/app/services"
	"github.com/T2-1c2023/NotificationsService/app/utilities"
)

func main() {
	notificationController := controller.New(&utilities.NotificationsSender{}, &services.UserService{})
	router := routes.SetupRouter(notificationController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002" // Default port
	}

	router.Run(":" + port)
}
