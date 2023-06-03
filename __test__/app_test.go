// __test__/app_test.go
package __test__

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	mock "github.com/T2-1c2023/NotificationsService/__mock__"
	"github.com/T2-1c2023/NotificationsService/app/controller"
	routes "github.com/T2-1c2023/NotificationsService/app/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetStatus(t *testing.T) {
	userServiceMock := mock.NewUserServiceMock()
	notificationSenderMock := mock.NewNotificationSenderMock()
	notificationController := controller.New(notificationSenderMock, userServiceMock)
	router := routes.SetupRouter(notificationController)

	// Create a test HTTP request to the / endpoint
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	// Create a test HTTP recorder
	recorder := httptest.NewRecorder()

	// Serve the request and record the response
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestNewFollowerNotification(t *testing.T) {
	userServiceMock := mock.NewUserServiceMock()
	notificationSenderMock := mock.NewNotificationSenderMock()
	notificationController := controller.New(notificationSenderMock, userServiceMock)
	router := routes.SetupRouter(notificationController)

	body := []byte(`{
		"followed_Id": 1
	}`)
	// Create a test HTTP request to the / endpoint
	req, _ := http.NewRequest(http.MethodPost, "/new-follower", bytes.NewBuffer(body))

	// Create a test HTTP recorder
	recorder := httptest.NewRecorder()

	// Serve the request and record the response
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 201, recorder.Code)
}
