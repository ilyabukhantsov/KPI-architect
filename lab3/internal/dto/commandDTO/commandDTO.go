package commandDTO

import "time"

type CreateTaskDTO struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type UpdateTaskDTO struct {
	Name               string  `json:"name" binding:"required"`
	UpdatedName        *string `json:"updated_name"`
	UpdatedDescription *string `json:"description"`
}

type DeleteTaskDTO struct {
	Name string `json:"name" binding:"required"`
}
