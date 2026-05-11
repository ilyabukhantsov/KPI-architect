package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"lab3/internal/config"
	"lab3/internal/domain"
	"lab3/internal/handler"
	"lab3/internal/model"
	"lab3/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFindAll_Integration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	config.DB = db
	config.DB.AutoMigrate(&model.Task{})

	config.DB.Create(&model.Task{Name: "Integration Task", Description: "Test Content"})

	repo := &repository.TaskRepository{}
	service := domain.NewTaskService(repo)
	queryH := handler.NewQueryHandler(service)
	cmdH := handler.NewCommandHandler(service)

	tc := NewTaskController(cmdH, queryH)

	r := gin.Default()
	r.GET("/FindAll", tc.FindAll)

	req, _ := http.NewRequest("GET", "/FindAll", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Integration Task")
}
