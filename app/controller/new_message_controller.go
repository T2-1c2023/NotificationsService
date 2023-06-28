package controller

import (
	"fmt"
	"net/http"

	"github.com/T2-1c2023/NotificationsService/app/model"
	"github.com/T2-1c2023/NotificationsService/app/services"
	"github.com/T2-1c2023/NotificationsService/app/utilities"
	"github.com/gin-gonic/gin"
)

type NewMessageController struct {
	Sender      utilities.INotificationSender
	UserService services.IUserService
	Logger      utilities.ILogger
}

type NewMessageInput struct {
	RecipientId int    `json:"recipient_id" binding:"required"`
	Message     string `json:"message" binding:"required"`
}

// NotifyNewMessage    		godoc
// @Summary      					Send new message notification to the given recipient user.
// @Description  					Send new message notification to the given recipient user.
// @Param				 					user_info header string true "Sender user decoded payload info"
// @Param									input body NewMessageInput true "Recipient user ID and message"
// @Success      					201
// @Failure								400
// @Failure								500
// @Router       					/new-message [post]
func (controller *NewFollowerController) NotifyNewMessage(c *gin.Context) {
	controller.Logger.LogInfo("GET /new-message")
	var requestBody NewMessageInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		controller.Logger.LogWarn("Bad request, returning 400")
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": err.Error()},
		)
		return
	}
	recipientUser, err := controller.UserService.GetUserById(requestBody.RecipientId, c.GetHeader("user_info"))
	if err != nil {
		controller.Logger.LogError(err)
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("User service error: %v\n", err)},
		)
		return
	}
	controller.Logger.LogDebug("Destinatario: " + recipientUser.Fullname)

	senderUser, err := model.NewUserFromUserInfo(c.GetHeader("user_info"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = controller.Sender.SendNotification(
		recipientUser.ExpoPushToken,
		"Nuevo mensaje de "+senderUser.Fullname,
		requestBody.Message)
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
