package handler

import (
	"lab3/internal/domain"
	"lab3/internal/dto/commandDTO"
	"lab3/internal/dto/queryDTO"
	"lab3/internal/mail"
	"lab3/internal/model"
)

type CommandHandler struct {
	service *domain.TaskService
}

func NewCommandHandler(s *domain.TaskService) *CommandHandler {
	return &CommandHandler{service: s}
}

func (h *CommandHandler) Create(dto commandDTO.CreateTaskDTO) error {
	task := &model.Task{Name: dto.Name, Description: dto.Description, Date: dto.Date}
	return h.service.Repo.Save(task)
}

func (h *CommandHandler) Update(dto commandDTO.UpdateTaskDTO) error {
	updates := make(map[string]interface{})

	if dto.UpdatedName != nil {
		updates["name"] = *dto.UpdatedName
	}
	if dto.UpdatedDescription != nil {
		updates["description"] = *dto.UpdatedDescription
	}

	return h.service.Repo.UpdateFields(dto.Name, updates)
}

func (h *CommandHandler) Delete(dto commandDTO.DeleteTaskDTO) error {
	return h.service.Repo.Delete(dto.Name)
}

type QueryHandler struct {
	service *domain.TaskService
}

func NewQueryHandler(s *domain.TaskService) *QueryHandler {
	return &QueryHandler{service: s}
}

func (h *QueryHandler) FindAll() ([]model.Task, error) {
	return h.service.Repo.GetAll()
}

func (h *QueryHandler) FindByName(dto queryDTO.FindByTaskDTO) (*model.Task, error) {
	return h.service.Repo.GetByName(*dto.Name)
}
