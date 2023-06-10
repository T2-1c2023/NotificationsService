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
	logFile, err := os.OpenFile("./log/app.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	logger := utilities.NewLogger(os.Getenv("LOG_LEVEL"), logFile)
	defer logger.CloseFile()
	logger.LogDebug("Microservice starting...")
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

	logger.LogInfo("App started on " + utilities.GetCurrentDate())
	logger.LogInfo("Notifications microservice running at port " + port)
	router.Run(":" + port)
}
