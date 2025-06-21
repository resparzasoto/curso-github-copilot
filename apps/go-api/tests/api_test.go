package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-api/src/controllers"
	"go-api/src/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users/", controllers.GetUsers)
	r.POST("/users/", controllers.CreateUser)
	r.GET("/health-check", controllers.HealthCheck)
	return r
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/users/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var users []models.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(users), 2)
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	user := models.User{Name: "Charlie"}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &createdUser)
	assert.NoError(t, err)
	assert.Equal(t, "Charlie", createdUser.Name)
}

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/health-check", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "running")
}
