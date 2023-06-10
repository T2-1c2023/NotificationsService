package __test__

import (
	"encoding/json"

	mock "github.com/T2-1c2023/NotificationsService/__mock__"
	"github.com/T2-1c2023/NotificationsService/app/controller"
	"github.com/T2-1c2023/NotificationsService/app/routes"
	"github.com/T2-1c2023/NotificationsService/app/utilities"
	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	notificationsSender := mock.NewNotificationSenderMock()
	userService := mock.NewUserServiceMock()
	logger := utilities.NewLogger("debug")
	notificationController := controller.NotificationController{
		Sender:      &notificationsSender,
		UserService: &userService,
		Logger:      &logger,
	}
	statusController := controller.StatusController{
		Logger: &logger,
	}
	router := routes.SetupRouter(&notificationController, &statusController)

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
