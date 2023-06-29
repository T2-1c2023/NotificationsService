package __test__

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/T2-1c2023/NotificationsService/__mock__"
	"github.com/T2-1c2023/NotificationsService/app/controller"
	"github.com/stretchr/testify/assert"
)

func TestNewMessageRequestReturns400WithoutUserInfo(t *testing.T) {
	userService := __mock__.NewUserServiceMock()
	notificationSender := __mock__.NewNotificationSenderMock()
	router := setUpRouter(&notificationSender, &userService)

	req, _ := http.NewRequest(http.MethodPost, "/notifications/new-message", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
	assert.Equal(t, notificationSender.Counter, 0)
}

func TestNewMessageRequestReturns400WithoutBody(t *testing.T) {
	userService := __mock__.NewUserServiceMock()
	notificationSender := __mock__.NewNotificationSenderMock()
	router := setUpRouter(&notificationSender, &userService)

	req, _ := http.NewRequest(http.MethodPost, "/notifications/new-follower", nil)
	recorder := httptest.NewRecorder()
	userInfo := getUserInfo(true)
	req.Header.Set("user_info", userInfo)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 400, recorder.Code)
	assert.Equal(t, notificationSender.Counter, 0)
}

func TestNewFollowerRequestReturns201WithCorrectBody(t *testing.T) {
	userService := __mock__.NewUserServiceMock()
	notificationSender := __mock__.NewNotificationSenderMock()
	router := setUpRouter(&notificationSender, &userService)

	requestBody := controller.NewMessageInput{
		RecipientId: 3,
		Message:     "Test message",
	}
	payload, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, "/notifications/new-message", bytes.NewBuffer(payload))
	recorder := httptest.NewRecorder()
	userInfo := getUserInfo(true)
	req.Header.Set("user_info", userInfo)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 201, recorder.Code)
	assert.Equal(t, notificationSender.Counter, 1)
}

func TestNewMessageRequestReturns500WhenUserServiceIsDown(t *testing.T) {
	userService := __mock__.NewErrorUserServiceMock()
	notificationSender := __mock__.NewNotificationSenderMock()
	router := setUpRouter(&notificationSender, &userService)

	requestBody := controller.NewMessageInput{
		RecipientId: 3,
		Message:     "Test message",
	}
	payload, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, "/notifications/new-message", bytes.NewBuffer(payload))
	recorder := httptest.NewRecorder()
	userInfo := getUserInfo(true)
	req.Header.Set("user_info", userInfo)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 500, recorder.Code)
	assert.Equal(t, notificationSender.Counter, 0)
}

func TestNewMessageRequestReturns500WhenNotificationWasntSent(t *testing.T) {
	userService := __mock__.NewUserServiceMock()
	notificationSender := __mock__.NewErrorNotificationSenderMock()
	router := setUpRouter(&notificationSender, &userService)

	requestBody := controller.NewMessageInput{
		RecipientId: 3,
		Message:     "Test message",
	}
	payload, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, "/notifications/new-message", bytes.NewBuffer(payload))
	recorder := httptest.NewRecorder()
	userInfo := getUserInfo(true)
	req.Header.Set("user_info", userInfo)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 500, recorder.Code)
}
