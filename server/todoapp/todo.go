package main

import (
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"encoding/json"
	"github.com/rs/cors"
)

var db, _ = gorm.Open("mysql", "root:root@/tododb?charset=utf8mb4&parseTime=True&loc=Local")

type TodoItem struct{
	Id int `gorm:"primary_key" json:"id"`
	Item string `sql:"type:VARCHAR(200)" json:"item"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	var ti TodoItem
   	decoder := json.NewDecoder(r.Body)
   	decoder.Decode(&ti)
	todo := &TodoItem{Item: ti.Item}
	db.Create(&todo)
	result := db.Last(&todo).Value
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(result)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var todos []TodoItem
	list := db.Find(&todos).Value
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(list)
}

func main() {
	defer db.Close()

	db.Debug().AutoMigrate(&TodoItem{})

	router := mux.NewRouter()
	router.HandleFunc("/todo", GetList).Methods("GET")
	router.HandleFunc("/todo", Add).Methods("POST")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	}).Handler(router)

	http.ListenAndServe(":3000", handler)
}