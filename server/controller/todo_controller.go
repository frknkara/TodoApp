package controller

import (
	"encoding/json"
	"net/http"
	"todoapp/entity"
	"todoapp/service"
)

type TodoController interface {
	Add(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service service.TodoService
}

func New(service service.TodoService) TodoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Add(w http.ResponseWriter, r *http.Request) {
	var ti entity.TodoItem
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&ti)
	todo := &entity.TodoItem{Item: ti.Item}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result, err := c.service.Add(*todo)
	if err != nil {
		http.Error(w, "Todo item couldn't be added", 400)
	} else {
		json.NewEncoder(w).Encode(result)
	}

}

func (c *controller) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result, err := c.service.GetList()
	if err != nil {
		http.Error(w, "List of todo items couldn't be fetched", 400)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}
