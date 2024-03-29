package controller

import (
	"fmt"
	"net/http"

	"github.com/T2-1c2023/NotificationsService/app/model"
	"github.com/T2-1c2023/NotificationsService/app/services"
	"github.com/T2-1c2023/NotificationsService/app/utilities"
	"github.com/gin-gonic/gin"
)

type NewFollowerController struct {
	Sender      utilities.INotificationSender
	UserService services.IUserService
	Logger      utilities.ILogger
}

type NewFollowerInput struct {
	FollowedId int `json:"followed_id" binding:"required"`
}

// NotifyNewFollower     	godoc
// @Summary      				 	Send new follower notification to the given followed user.
// @Description  					Send new follower notification to the given followed user.
// @Param				 					user_info header string true "Follower user decoded payload info"
// @Param									input body NewFollowerInput true "Followed user ID"
// @Success      					201
// @Failure								400
// @Failure			 					500
// @Router       					/new-follower [post]
func (controller *NewFollowerController) NotifyNewFollower(c *gin.Context) {
	controller.Logger.LogInfo("GET /new-follower")
	var requestBody NewFollowerInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		controller.Logger.LogWarn("Bad request, returning 400")
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": err.Error()},
		)
		return
	}
	followedUser, err := controller.UserService.GetUserById(requestBody.FollowedId, c.GetHeader("user_info"))
	if err != nil {
		controller.Logger.LogError(err)
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("User service error: %v\n", err)},
		)
		return
	}

	followerUser, err := model.NewUserFromUserInfo(c.GetHeader("user_info"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	controller.Logger.LogDebug("Usuario seguido: " + followedUser.Fullname)

	err = controller.Sender.SendNotification(
		followedUser.ExpoPushToken,
		"Nuevo seguidor!",
		"El usuario "+followerUser.Fullname+" te empezó a seguir.")
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
