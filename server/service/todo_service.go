package service

import (
	"errors"
	"todoapp/dcl"
	"todoapp/entity"
)

type TodoService interface {
	Add(entity.TodoItem) (entity.TodoItem, error)
	GetList() ([]entity.TodoItem, error)
}

type todoService struct {
	dcl dcl.TodoDcl
}

func New(dcl dcl.TodoDcl) TodoService {
	return &todoService{
		dcl: dcl,
	}
}

func (service *todoService) Add(todoItem entity.TodoItem) (entity.TodoItem, error) {
	if todoItem.Item == "" {
		return entity.TodoItem{}, errors.New("item can't be empty")
	}
	if len(todoItem.Item) > 200 {
		return entity.TodoItem{}, errors.New("item length can't be longer than 200")
	}
	return service.dcl.Add(todoItem)
}

func (service *todoService) GetList() ([]entity.TodoItem, error) {
	return service.dcl.GetList()
}
