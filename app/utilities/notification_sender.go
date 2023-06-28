package utilities

import (
	"fmt"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type INotificationSender interface {
	SendNotification(pushTokenString string, title string, body string) error
}

type NotificationSender struct {
	Logger ILogger
}

func (sender *NotificationSender) SendNotification(pushTokenString string, title string, body string) error {
	sender.Logger.LogDebug("Expo push token: " + pushTokenString)
	var pushToken, err = expo.NewExponentPushToken(pushTokenString)
	if err != nil {
		return err
	}

	client := expo.NewPushClient(nil)

	response, err := client.Publish(
		&expo.PushMessage{
			To:       []expo.ExponentPushToken{pushToken},
			Body:     body,
			Sound:    "default",
			Title:    title,
			Priority: expo.DefaultPriority,
		},
	)
	if err != nil {
		return fmt.Errorf("Error: " + err.Error())
	}
	sender.Logger.LogDebug("Response ID: " + response.ID)
	sender.Logger.LogDebug("Response Status: " + response.Status)
	sender.Logger.LogDebug("Response Sound: " + response.PushMessage.Sound)
	if response.ValidateResponse() != nil {
		return fmt.Errorf("couldn't validate push notification response")
	}
	return nil
}
