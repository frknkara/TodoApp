package entity

type TodoItem struct {
	Id   int    `gorm:"primary_key" json:"id"`
	Item string `sql:"type:VARCHAR(200)" json:"item"`
}
