package repository

import (
	"lab3/internal/config"
	"lab3/internal/model"
)

type TaskRepository struct{}

func (r *TaskRepository) Save(task *model.Task) error {
	return config.DB.Create(task).Error
}

func (r *TaskRepository) Update(task *model.Task) error {
	return config.DB.Model(&model.Task{}).Where("name = ?", task.Name).Updates(task).Error
}

func (r *TaskRepository) UpdateFields(name string, fields map[string]interface{}) error {
	return config.DB.Model(&model.Task{}).
		Where("name = ?", name).
		Updates(fields).Error
}

func (r *TaskRepository) Delete(name string) error {
	return config.DB.Where("name = ?", name).Delete(&model.Task{}).Error
}

func (r *TaskRepository) GetByName(name string) (*model.Task, error) {
	var task model.Task
	err := config.DB.Where("name = ?", name).First(&task).Error
	return &task, err
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	err := config.DB.Find(&tasks).Error
	return tasks, err
}
