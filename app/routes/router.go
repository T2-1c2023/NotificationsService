package routes

import (
	controller "github.com/T2-1c2023/NotificationsService/app/controller"
	"github.com/T2-1c2023/NotificationsService/app/validation"
	_ "github.com/T2-1c2023/NotificationsService/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(
	newFollowerController *controller.NewFollowerController,
	newMessageController *controller.NewMessageController,
	trainingCompletedController *controller.TrainingCompletedController,
	statusController *controller.StatusController) *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/", statusController.GetStatus)

	router.GET("/health", statusController.GetHealth)

	router.GET("/status",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		statusController.GetServiceStatus)
	router.POST("/status",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		statusController.ChangeServiceStatus)

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/logs", statusController.GetLogs)

	router.Use(statusController.ValidateBlockedStatus)

	router.POST("/notifications/new-follower",
		validation.UserInfoHeaderValidator,
		newFollowerController.NotifyNewFollower)

	router.POST("/notifications/new-message",
		validation.UserInfoHeaderValidator,
		newFollowerController.NotifyNewMessage)

	router.POST("/notifications/training-completed",
		validation.UserInfoHeaderValidator,
		trainingCompletedController.NotifyTrainingCompleted)

	return router
}
