package main

import (
	"lab3/internal/config"
	"lab3/internal/controller"
	"lab3/internal/domain"
	"lab3/internal/handler"
	"lab3/internal/mailnotifer"
	"lab3/internal/model"
	"lab3/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.Connect()
	config.DB.AutoMigrate(&model.Task{})

	repo := &repository.TaskRepository{}
	service := domain.NewTaskService(repo)

	cmdHandler := handler.NewCommandHandler(service)
	queryHandler := handler.NewQueryHandler(service)
	mailService := mailnotifer.NewMailService()

	tc := controller.NewTaskController(cmdHandler, queryHandler, mailService)

	r.POST("/CreateTask", tc.CreateTask)
	r.PATCH("/UpdateTask", tc.UpdateTask)
	r.DELETE("/DeleteTask", tc.DeleteTask)
	r.GET("/FindByName", tc.FindByName)
	r.GET("/FindAll", tc.FindAll)

	r.Run(":8080")
}
