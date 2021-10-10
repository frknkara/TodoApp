package data

import (
	"log"
	"todoapp/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDb() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("mysql", dbDSN)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//db.Debug().DropTableIfExists(&entity.TodoItem{})
	db.Debug().AutoMigrate(&entity.TodoItem{})
	return db, nil
}

func CloseConnection(db *gorm.DB) error {
	if err := db.Close(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
