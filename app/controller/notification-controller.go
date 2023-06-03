package controller

import (
	"fmt"
	"net/http"

	services "github.com/T2-1c2023/NotificationsService/app/services"
	utilities "github.com/T2-1c2023/NotificationsService/app/utilities"
	"github.com/gin-gonic/gin"
)

type NewFollowerInput struct {
	Followed_ID int `json:"followed_id" binding:"required"`
}

// GetStatus     godoc
// @Summary      Send new follower notification to the given followed user.
// @Description  Send new follower notification to the given followed user.
// @Success      201
// @Router       / [post]
func NotifyNewFollower(c *gin.Context) {
	var requestBody NewFollowerInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": err.Error()},
		)
		return
	}
	followedUser, err := services.GetUserById(requestBody.Followed_ID, c.GetHeader("user_info"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("User service error: %v\n", err)},
		)
		return
	}

	err = utilities.SendNotification(
		followedUser.Expo_Push_Token,
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
