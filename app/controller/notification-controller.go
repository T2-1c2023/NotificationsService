package controller

import (
	"fmt"
	"net/http"

	"github.com/T2-1c2023/NotificationsService/app/model"
	"github.com/gin-gonic/gin"
)

type (
	sender interface {
		SendNotification(pushTokenString string, title string, body string) error
	}
	userService interface {
		GetUserById(id int, userInfo string) (model.User, error)
	}
	NotificationController struct {
		sender      sender
		userService userService
	}
)

func New(sender sender, userService userService) *NotificationController {
	return &NotificationController{
		sender:      sender,
		userService: userService,
	}
}

type NewFollowerInput struct {
	FollowedId int `json:"followed_id" binding:"required"`
}

// GetStatus     godoc
// @Summary      Send new follower notification to the given followed user.
// @Description  Send new follower notification to the given followed user.
// @Success      201
// @Router       / [post]
func (controller *NotificationController) NotifyNewFollower(c *gin.Context) {
	var requestBody NewFollowerInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": err.Error()},
		)
		return
	}
	followedUser, err := controller.userService.GetUserById(requestBody.FollowedId, c.GetHeader("user_info"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("User service error: %v\n", err)},
		)
		return
	}

	err = controller.sender.SendNotification(
		followedUser.ExpoPushToken,
		"Nuevo seguidor!",
		"El usuario "+followedUser.Fullname+" te empezó a seguir.")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Utilities error: %v", err.Error())},
		)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Notification sent",
	})
}
