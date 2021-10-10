package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapp/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTodoService struct {
	mock.Mock
}

func (mock *MockTodoService) Add(todoItem entity.TodoItem) (entity.TodoItem, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.TodoItem), args.Error(1)
}

func (mock *MockTodoService) GetList() ([]entity.TodoItem, error) {
	args := mock.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]entity.TodoItem), args.Error(1)
}

func TestAdd(t *testing.T) {
	mockService := new(MockTodoService)

	todoItem := entity.TodoItem{Id: 1, Item: "TestItem"}
	mockService.On("Add").Return(todoItem, nil)

	var jsonBody = []byte(`{item: "TestItem"}`)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonBody))

	controller := New(mockService)
	handler := http.HandlerFunc(controller.Add)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Invalid status code. Expected: %v but got %v", http.StatusOK, status)
	}

	mockService.AssertExpectations(t)

	var result entity.TodoItem
	json.NewDecoder(io.Reader(response.Body)).Decode(&result)

	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "TestItem", result.Item)
}

func TestAddError(t *testing.T) {
	mockService := new(MockTodoService)

	mockService.On("Add").Return(entity.TodoItem{}, errors.New("Add test error"))

	var jsonBody = []byte(`{item: "TestItem"}`)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonBody))

	controller := New(mockService)
	handler := http.HandlerFunc(controller.Add)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusBadRequest {
		t.Errorf("Invalid status code. Expected: %v but got %v", http.StatusBadRequest, status)
	}

	mockService.AssertExpectations(t)

	result, _ := io.ReadAll(response.Body)

	assert.Contains(t, string(result), "Todo item couldn't be added")
}

func TestGetList(t *testing.T) {
	mockService := new(MockTodoService)

	todoItem := entity.TodoItem{Id: 1, Item: "TestItem"}
	mockService.On("GetList").Return([]entity.TodoItem{todoItem}, nil)

	req, _ := http.NewRequest("GET", "/todo", nil)

	controller := New(mockService)
	handler := http.HandlerFunc(controller.GetList)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Invalid status code. Expected: %v but got %v", http.StatusOK, status)
	}

	mockService.AssertExpectations(t)

	var list []entity.TodoItem
	json.NewDecoder(io.Reader(response.Body)).Decode(&list)

	assert.NotNil(t, list)
	assert.Equal(t, 1, list[0].Id)
	assert.Equal(t, "TestItem", list[0].Item)
}

func TestGetListError(t *testing.T) {
	mockService := new(MockTodoService)

	mockService.On("GetList").Return(nil, errors.New("GetList test error"))

	req, _ := http.NewRequest("GET", "/todo", nil)

	controller := New(mockService)
	handler := http.HandlerFunc(controller.GetList)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusBadRequest {
		t.Errorf("Invalid status code. Expected: %v but got %v", http.StatusBadRequest, status)
	}

	mockService.AssertExpectations(t)

	result, _ := io.ReadAll(response.Body)

	assert.Contains(t, string(result), "List of todo items couldn't be fetched")
}
