package __mock__

type NotificationsSenderMock struct{}

func NewNotificationSenderMock() NotificationsSenderMock {
	return NotificationsSenderMock{}
}

func (sender *NotificationsSenderMock) SendNotification(pushTokenString string, title string, body string) error {
	return nil
}
