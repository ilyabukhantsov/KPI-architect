package handler

import (
	"lab3/internal/domain"
	"lab3/internal/dto/commandDTO"
	"lab3/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct{ mock.Mock }

func (m *MockRepo) Save(t *model.Task) error   { return m.Called(t).Error(0) }
func (m *MockRepo) Update(t *model.Task) error { return m.Called(t).Error(0) }
func (m *MockRepo) Delete(n string) error      { return m.Called(n).Error(0) }
func (m *MockRepo) UpdateFields(n string, f map[string]interface{}) error {
	return m.Called(n, f).Error(0)
}
func (m *MockRepo) GetAll() ([]model.Task, error)           { return nil, nil }
func (m *MockRepo) GetByName(n string) (*model.Task, error) { return nil, nil }

func TestCreateTask_Invariant(t *testing.T) {
	mockRepo := new(MockRepo)
	service := domain.NewTaskService(mockRepo)
	h := NewCommandHandler(service)

	dto := commandDTO.CreateTaskDTO{Name: "Clean Code"}

	mockRepo.On("Save", mock.MatchedBy(func(t *model.Task) bool {
		return t.Name == "Clean Code"
	})).Return(nil).Once()

	err := h.Create(dto)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
