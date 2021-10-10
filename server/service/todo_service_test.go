package service

import (
	"errors"
	"testing"
	"todoapp/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTodoDcl struct {
	mock.Mock
}

func (mock *MockTodoDcl) InitDb() error {
	return nil
}

func (mock *MockTodoDcl) CloseConnection() error {
	return nil
}

func (mock *MockTodoDcl) Add(entity.TodoItem) (entity.TodoItem, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.TodoItem), args.Error(1)
}

func (mock *MockTodoDcl) GetList() ([]entity.TodoItem, error) {
	args := mock.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]entity.TodoItem), args.Error(1)
}

func TestGetList(t *testing.T) {
	mockDcl := new(MockTodoDcl)

	todoItem := entity.TodoItem{Id: 1, Item: "TestItem"}
	mockDcl.On("GetList").Return([]entity.TodoItem{todoItem}, nil)

	service := New(mockDcl)

	result, _ := service.GetList()

	mockDcl.AssertExpectations(t)

	assert.Equal(t, 1, result[0].Id)
	assert.Equal(t, "TestItem", result[0].Item)
}

func TestGetListError(t *testing.T) {
	mockDcl := new(MockTodoDcl)

	mockDcl.On("GetList").Return(nil, errors.New("GetList test error"))

	service := New(mockDcl)

	result, err := service.GetList()

	mockDcl.AssertExpectations(t)

	assert.Nil(t, result)
	assert.Equal(t, "GetList test error", err.Error())
}

func TestAdd(t *testing.T) {
	mockDcl := new(MockTodoDcl)

	todoItem := entity.TodoItem{Item: "TestItem"}
	mockDcl.On("Add").Return(todoItem, nil)

	service := New(mockDcl)

	result, _ := service.Add(todoItem)

	mockDcl.AssertExpectations(t)

	assert.Equal(t, "TestItem", result.Item)
}

func TestAddError(t *testing.T) {
	mockDcl := new(MockTodoDcl)

	todoItem := entity.TodoItem{Item: "TestItem"}
	mockDcl.On("Add").Return(entity.TodoItem{}, errors.New("Add test error"))

	service := New(mockDcl)

	_, err := service.Add(todoItem)

	mockDcl.AssertExpectations(t)

	assert.Equal(t, "Add test error", err.Error())
}

func TestAddEmptyTodoItem(t *testing.T) {
	mockDcl := new(MockTodoDcl)

	service := New(mockDcl)

	result, err := service.Add(entity.TodoItem{})

	assert.Equal(t, entity.TodoItem{}, result)
	assert.Equal(t, "item can't be empty", err.Error())
}

func TestAddEmptyTextTodoItem(t *testing.T) {
	mockDcl := new(MockTodoDcl)

	todoItem := entity.TodoItem{Item: ""}

	service := New(mockDcl)

	result, err := service.Add(todoItem)

	assert.Equal(t, entity.TodoItem{}, result)
	assert.Equal(t, "item can't be empty", err.Error())
}

func TestAddLongerThan200BytesTextTodoItem(t *testing.T) {
	mockDcl := new(MockTodoDcl)

	sample201BytesText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque mi elit, tincidunt nec pharetra a, pellentesque sed orci. Cras et eleifend massa. Suspendisse sit amet aliquet orci nullam sodales."
	todoItem := entity.TodoItem{Item: sample201BytesText}

	service := New(mockDcl)

	result, err := service.Add(todoItem)

	assert.Equal(t, entity.TodoItem{}, result)
	assert.Equal(t, "item length can't be longer than 200", err.Error())
}
