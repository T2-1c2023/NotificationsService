// __test__/app_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	routes "github.com/T2-1c2023/NotificationsService/app/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetStatus(t *testing.T) {
	router := routes.SetupRouter()

	// Create a test HTTP request to the / endpoint
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	// Create a test HTTP recorder
	recorder := httptest.NewRecorder()

	// Serve the request and record the response
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}
