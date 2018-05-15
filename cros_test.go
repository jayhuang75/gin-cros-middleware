package cros

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORSMiddleware(t *testing.T) {

	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.Use(CORSMiddleware)
	r.GET("/api")

	////////////////////////////////////
	// Test with header
	////////////////////////////////////
	req1, _ := http.NewRequest(http.MethodGet, "/api", nil)
	req1.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	// Create a response recorder so you can inspect the response
	resp1 := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(resp1, req1)

	assert.Equal(t, resp1.Code, 200)

	assert.Equal(t, req1.Header.Get("Access-Control-Allow-Methods"), "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

	////////////////////////////////////
	// Test header with options
	////////////////////////////////////
	req2, _ := http.NewRequest(http.MethodGet, "/api", nil)
	req2.Method = "OPTIONS"
	// Create a response recorder so you can inspect the response
	resp2 := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(resp2, req2)

	assert.Equal(t, resp2.Code, 200)

	// assert.Equal(t, req1.Header.Get("Access-Control-Allow-Methods"), "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

}
