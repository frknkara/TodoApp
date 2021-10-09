package dcl

import (
	"log"
	"todoapp/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TodoDcl interface {
	InitDb() error
	CloseConnection() error
	Add(entity.TodoItem) (entity.TodoItem, error)
	GetList() ([]entity.TodoItem, error)
}

type todoDcl struct {
	db *gorm.DB
}

func New() TodoDcl {
	return &todoDcl{}
}

func (dcl *todoDcl) InitDb() error {
	var err error
	dcl.db, err = gorm.Open("mysql", dbDSN)
	if err != nil {
		log.Println(err)
		return err
	}
	dcl.db.Debug().DropTableIfExists(&entity.TodoItem{})
	dcl.db.Debug().AutoMigrate(&entity.TodoItem{})
	return nil
}

func (dcl *todoDcl) CloseConnection() error {
	if err := dcl.db.Close(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (dcl *todoDcl) Add(todoItem entity.TodoItem) (entity.TodoItem, error) {
	err := dcl.db.Create(&todoItem).Error
	if err != nil {
		log.Println(err)
		return entity.TodoItem{}, err
	}
	return todoItem, nil
}

func (dcl *todoDcl) GetList() ([]entity.TodoItem, error) {
	var todos []entity.TodoItem
	err := dcl.db.Find(&todos).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todos, nil
}
