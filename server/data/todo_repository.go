package data

import (
	"log"
	"todoapp/entity"

	"github.com/jinzhu/gorm"
)

type TodoRepository interface {
	Add(entity.TodoItem) (entity.TodoItem, error)
	GetList() ([]entity.TodoItem, error)
}

type todoRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (repository *todoRepository) Add(todoItem entity.TodoItem) (entity.TodoItem, error) {
	err := repository.db.Create(&todoItem).Error
	if err != nil {
		log.Println(err)
		return entity.TodoItem{}, err
	}
	return todoItem, nil
}

func (repository *todoRepository) GetList() ([]entity.TodoItem, error) {
	var todos []entity.TodoItem
	err := repository.db.Find(&todos).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todos, nil
}
