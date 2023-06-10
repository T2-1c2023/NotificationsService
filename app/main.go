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
	logger := utilities.NewLogger(os.Getenv("LOG_LEVEL"))
	notificationController := controller.NotificationController{
		Sender:      &utilities.NotificationsSender{},
		UserService: &services.UserService{},
	}
	statusController := controller.StatusController{
		Logger: &logger,
	}
	router := routes.SetupRouter(&notificationController, &statusController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002" // Default port
	}

	logger.LogInfo("Aplicación de Golang ejecutándose en puerto " + port)
	router.Run(":" + port)
}
