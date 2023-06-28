package __mock__

import "fmt"

type NotificationSenderMock struct {
	Counter int
}

func NewNotificationSenderMock() NotificationSenderMock {
	return NotificationSenderMock{}
}

func (sender *NotificationSenderMock) SendNotification(pushTokenString string, title string, body string) error {
	sender.Counter++
	return nil
}

type ErrorNotificationSenderMock struct{}

func NewErrorNotificationSenderMock() ErrorNotificationSenderMock {
	return ErrorNotificationSenderMock{}
}

func (sender *ErrorNotificationSenderMock) SendNotification(pushTokenString string, title string, body string) error {
	return fmt.Errorf("error in ErrorNotificationsSenderMock")
}
