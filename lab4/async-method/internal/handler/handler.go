package handler

import (
	"lab3/internal/domain"
	"lab3/internal/dto/commandDTO"
	"lab3/internal/dto/queryDTO"
	"lab3/internal/events"
	"lab3/internal/model"
)

type CommandHandler struct {
	service *domain.TaskService
	bus     *events.EventBus
}

func NewCommandHandler(s *domain.TaskService, b *events.EventBus) *CommandHandler {
	return &CommandHandler{service: s, bus: b}
}

func (h *CommandHandler) Create(dto commandDTO.CreateTaskDTO) error {
	task := &model.Task{Name: dto.Name, Description: dto.Description, Date: dto.Date}
	if err := h.service.Repo.Save(task); err != nil {
		return err
	}
	h.bus.Publish("task.created", dto)
	return nil
}

func (h *CommandHandler) Update(dto commandDTO.UpdateTaskDTO) error {
	updates := make(map[string]interface{})
	if dto.UpdatedName != nil {
		updates["name"] = *dto.UpdatedName
	}
	if dto.UpdatedDescription != nil {
		updates["description"] = *dto.UpdatedDescription
	}

	if err := h.service.Repo.UpdateFields(dto.Name, updates); err != nil {
		return err
	}
	h.bus.Publish("task.updated", dto)
	return nil
}

func (h *CommandHandler) Delete(dto commandDTO.DeleteTaskDTO) error {
	if err := h.service.Repo.Delete(dto.Name); err != nil {
		return err
	}
	h.bus.Publish("task.deleted", dto)
	return nil
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
