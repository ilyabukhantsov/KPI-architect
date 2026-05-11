package main

import (
	"fmt"
	"lab3/internal/config"
	"lab3/internal/controller"
	"lab3/internal/domain"
	"lab3/internal/events"
	"lab3/internal/handler"
	"lab3/internal/mailnotifer"
	"lab3/internal/model"
	"lab3/internal/phonenotifer"
	"lab3/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.Connect()
	config.DB.AutoMigrate(&model.Task{})

	bus := events.NewEventBus()

	mailService := mailnotifer.NewMailService()
	phoneService := phonenotifer.NewPhoneService()

	go func() {
		ch := bus.Subscribe()
		for event := range ch {
			fmt.Printf("Logging event: %s\n", event.Name)
			mailService.MailSomeone("Action: " + event.Name)
			phoneService.PhoneClient()
		}
	}()

	repo := &repository.TaskRepository{}
	service := domain.NewTaskService(repo)

	cmdHandler := handler.NewCommandHandler(service, bus)
	queryHandler := handler.NewQueryHandler(service)

	tc := controller.NewTaskController(cmdHandler, queryHandler)

	r.POST("/CreateTask", tc.CreateTask)
	r.PATCH("/UpdateTask", tc.UpdateTask)
	r.DELETE("/DeleteTask", tc.DeleteTask)
	r.GET("/FindByName", tc.FindByName)

	r.Run(":8080")
}
