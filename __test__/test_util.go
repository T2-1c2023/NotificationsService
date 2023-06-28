package __test__

import (
	"encoding/json"
	"os"

	"github.com/T2-1c2023/NotificationsService/app/controller"
	"github.com/T2-1c2023/NotificationsService/app/routes"
	"github.com/T2-1c2023/NotificationsService/app/services"
	"github.com/T2-1c2023/NotificationsService/app/utilities"
	"github.com/gin-gonic/gin"
)

func setUpRouter(
	notificationsSender utilities.INotificationSender,
	userService services.IUserService) *gin.Engine {
	logFile, _ := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger := utilities.NewLogger("debug", logFile)
	newFollowerController := controller.NewFollowerController{
		Sender:      notificationsSender,
		UserService: userService,
		Logger:      &logger,
	}
	statusController := controller.StatusController{
		Logger: &logger,
	}
	router := routes.SetupRouter(&newFollowerController, &statusController)

	return router
}

func getUserInfo(isAdmin bool) string {
	userInfo, _ := json.Marshal(
		struct {
			Id      int  `json:"id"`
			IsAdmin bool `json:"is_admin"`
		}{
			Id:      1,
			IsAdmin: isAdmin,
		},
	)
	return string(userInfo)
}
