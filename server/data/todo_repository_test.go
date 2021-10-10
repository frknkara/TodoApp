package data

import (
	"regexp"
	"testing"
	"todoapp/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open("mysql", db)
	return gdb, mock
}

func TestGetList(t *testing.T) {
	db, mock := setup(t)
	defer db.Close()

	repository := New(db)

	query := "SELECT * FROM `todo_items`"

	rows := sqlmock.
		NewRows([]string{"id", "item"}).
		AddRow(1, "Test Item 1").
		AddRow(2, "Test Item 2")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	list, err := repository.GetList()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(list))
	assert.Equal(t, 1, list[0].Id)
	assert.Equal(t, 2, list[1].Id)
	assert.Equal(t, "Test Item 1", list[0].Item)
	assert.Equal(t, "Test Item 2", list[1].Item)
}

func TestGetListEmpty(t *testing.T) {
	db, mock := setup(t)
	defer db.Close()

	repository := New(db)

	query := "SELECT * FROM `todo_items`"

	rows := sqlmock.
		NewRows([]string{"id", "item"})

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	list, err := repository.GetList()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(list))
}

func TestAdd(t *testing.T) {
	db, mock := setup(t)
	defer db.Close()

	repository := New(db)

	query := "INSERT INTO `todo_items` (`item`) VALUES (?)"

	todoItem := entity.TodoItem{Item: "Test Item"}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(todoItem.Item).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := repository.Add(todoItem)
	assert.Nil(t, err)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Test Item", result.Item)
}
