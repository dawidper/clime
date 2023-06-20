package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/:action", handleRoute)
	return router
}

func TestHandleRoute(t *testing.T) {
	router := setupRouter()

	// Test addition
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/add?x=3&y=5", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"action":"add","x":3,"y":5,"answer":8,"cached":false}`, w.Body.String())

	// Test subtraction
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/subtract?x=10&y=4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"action":"subtract","x":10,"y":4,"answer":6,"cached":false}`, w.Body.String())

	// Test multiplication
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/multiply?x=5&y=5", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"action":"multiply","x":5,"y":5,"answer":25,"cached":false}`, w.Body.String())

	// Test division
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/divide?x=10&y=2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"action":"divide","x":10,"y":2,"answer":5,"cached":false}`, w.Body.String())

	// Test division by zero
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/divide?x=10&y=0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, `"`+errorDivisionByZero+`"`, w.Body.String())

	// Test invalid action
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/invalidAction?x=10&y=2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `"`+errorMessageWrongAction+`"`, w.Body.String())

	// Test invalid parameters
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/add?x=abc&y=2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `"`+errorMessageWrongParameter+`"`, w.Body.String())
}
