package core

import (
	"github.com/renantarouco/clean-todo/model"
)

type TasksRepository interface {
	Put(model.Task) error
	All() ([]model.Task, error)
}

type TasksService struct {
	tasksRepo   TasksRepository
	addCounter  int
	listCounter int
}

func NewTasksService(tasksRepo TasksRepository) *TasksService {
	return &TasksService{
		tasksRepo:   tasksRepo,
		addCounter:  0,
		listCounter: 0,
	}
}

func (s *TasksService) Add(task model.Task) error {
	if err := s.tasksRepo.Put(task); err != nil {
		return err
	}
	s.addCounter += 1
	return nil
}

func (s *TasksService) List() ([]model.Task, error) {
	tasks, err := s.tasksRepo.All()
	if err != nil {
		return nil, err
	}
	s.listCounter += 1
	return tasks, nil
}
