package controller

import (
	"fmt"
	"net/http"

	"github.com/T2-1c2023/NotificationsService/app/model"
	"github.com/T2-1c2023/NotificationsService/app/services"
	"github.com/T2-1c2023/NotificationsService/app/utilities"
	"github.com/gin-gonic/gin"
)

type TrainingCompletedController struct {
	Sender      utilities.INotificationSender
	UserService services.IUserService
	Logger      utilities.ILogger
}

type TrainingCompletedInput struct {
	TrainingTitle string `json:"training_title" binding:"required"`
}

// NotifyTrainingCompleted    		godoc
// @Summary      									Send completed training notification to the user.
// @Description  									Send completed training notification to the user.
// @Param				 									user_info header string true "User decoded payload info"
// @Param													input body TrainingCompletedInput true "Completed training title"
// @Success      									201
// @Failure												400
// @Failure												500
// @Router       									/training-completed [post]
func (controller *TrainingCompletedController) NotifyTrainingCompleted(c *gin.Context) {
	controller.Logger.LogInfo("GET /training-completed")
	var requestBody TrainingCompletedInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		controller.Logger.LogWarn("Bad request, returning 400")
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": err.Error()},
		)
		return
	}

	user, err := model.NewUserFromUserInfo(c.GetHeader("user_info"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	controller.Logger.LogDebug("Usuario: " + user.Fullname)

	err = controller.Sender.SendNotification(
		user.ExpoPushToken,
		"Entrenamiento completado",
		"Felicitaciones, "+user.Fullname+", completaste el entrenamiento "+requestBody.TrainingTitle)
	if err != nil {
		controller.Logger.LogError(err)
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Utilities error: %v", err.Error())},
		)
		return
	}
	controller.Logger.LogInfo("Notification sent")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Notification sent",
	})
}
