package domain

import "lab3/internal/model"

type TaskRepository interface {
	Save(task *model.Task) error
	Update(task *model.Task) error
	Delete(name string) error
	GetByName(name string) (*model.Task, error)
	GetAll() ([]model.Task, error)
	UpdateFields(name string, fields map[string]interface{}) error
}

type TaskService struct {
	Repo TaskRepository
}

func NewTaskService(r TaskRepository) *TaskService {
	return &TaskService{Repo: r}
}
