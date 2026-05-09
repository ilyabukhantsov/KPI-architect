package controller

import (
	"github.com/gin-gonic/gin"
	"lab3/internal/dto/commandDTO"
	"lab3/internal/dto/queryDTO"
	"lab3/internal/handler"
	"lab3/internal/mailnotifer"
	"lab3/internal/phonenotifer"
)

type TaskController struct {
	Command *handler.CommandHandler
	Query   *handler.QueryHandler
	Mail    *mailnotifer.Mail
	Phone   *phonenotifer.PhoneService
}

func NewTaskController(c *handler.CommandHandler, q *handler.QueryHandler, m *mailnotifer.Mail, p *phonenotifer.PhoneService) *TaskController {
	return &TaskController{Command: c, Query: q, Mail: m, Phone: p}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var dto commandDTO.CreateTaskDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := tc.Command.Create(dto); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	tc.Mail.MailSomeone("All good")
	tc.Phone.PhoneClient()
	c.JSON(201, gin.H{"message": "Created"})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	var dto commandDTO.UpdateTaskDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := tc.Command.Update(dto); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	tc.Mail.MailSomeone("All good")
	tc.Phone.PhoneClient()

	c.JSON(200, gin.H{"message": "Updated"})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	var dto commandDTO.DeleteTaskDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := tc.Command.Delete(dto); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	tc.Mail.MailSomeone("All good")
	tc.Phone.PhoneClient()
	c.JSON(200, gin.H{"message": "Deleted"})
}

func (tc *TaskController) FindByName(c *gin.Context) {
	var dto queryDTO.FindByTaskDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	task, err := tc.Query.FindByName(dto)
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	tc.Mail.MailSomeone("All good")
	tc.Phone.PhoneClient()
	c.JSON(200, task)
}

func (tc *TaskController) FindAll(c *gin.Context) {
	tasks, _ := tc.Query.FindAll()
	tc.Mail.MailSomeone("All good")
	tc.Phone.PhoneClient()
	c.JSON(200, tasks)
}
