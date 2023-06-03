package routes

import (
	controller "github.com/T2-1c2023/NotificationsService/app/controller"
	_ "github.com/T2-1c2023/NotificationsService/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(notificationController *controller.NotificationController) *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	router.GET("/", controller.GetStatus)

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", controller.GetHealth)

	router.POST("/new-follower", notificationController.NotifyNewFollower)

	return router
}
