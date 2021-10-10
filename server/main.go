package main

import (
	"log"
	"todoapp/controller"
	"todoapp/data"
	"todoapp/router"
	"todoapp/service"

	"net/http"

	"github.com/rs/cors"
)

func main() {
	db, err := data.InitDb()
	if err != nil {
		log.Fatalln(err)
	}
	var (
		todoRepository data.TodoRepository       = data.New(db)
		todoService    service.TodoService       = service.New(todoRepository)
		todoController controller.TodoController = controller.New(todoService)
	)

	defer data.CloseConnection(db)
	r := router.Router(todoController)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	}).Handler(r)

	port := ":3000"

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, handler))
}
