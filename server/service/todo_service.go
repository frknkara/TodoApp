package service

import (
	"errors"
	"todoapp/data"
	"todoapp/entity"
)

type TodoService interface {
	Add(entity.TodoItem) (entity.TodoItem, error)
	GetList() ([]entity.TodoItem, error)
}

type todoService struct {
	repository data.TodoRepository
}

func New(repository data.TodoRepository) TodoService {
	return &todoService{
		repository: repository,
	}
}

func (service *todoService) Add(todoItem entity.TodoItem) (entity.TodoItem, error) {
	if todoItem.Item == "" {
		return entity.TodoItem{}, errors.New("item can't be empty")
	}
	if len(todoItem.Item) > 200 {
		return entity.TodoItem{}, errors.New("item length can't be longer than 200")
	}
	return service.repository.Add(todoItem)
}

func (service *todoService) GetList() ([]entity.TodoItem, error) {
	return service.repository.GetList()
}
