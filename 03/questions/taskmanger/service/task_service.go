package service

import "github.com/DKeshavarz/goraz/module3/taskmanger/model"

type taskRepository interface {
	FetchAll() ([]model.Task, error)
	Create(task model.Task) error
	GetByID(id int) (*model.Task, error)
	Delete(id int) error
}

type TaskService interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task model.Task) error
	// ...
}

type taskService struct {
	repo taskRepository
}

func NewTaskService(repo taskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) GetAllTasks() ([]model.Task, error) {
	// TODO: Call the repository
	return nil, nil
}

func (s *taskService) CreateTask(task model.Task) error{
	return nil
}
