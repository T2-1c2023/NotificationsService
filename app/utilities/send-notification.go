package utilities

import (
	"fmt"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

func SendNotification(pushTokenString string, title string, body string) error {
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
	if response.ValidateResponse() != nil {
		return fmt.Errorf("couldn't validate push notification response")
	}
	return nil
}
