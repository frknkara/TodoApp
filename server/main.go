package main

import (
	"log"
	"todoapp/controller"
	"todoapp/dcl"
	"todoapp/router"
	"todoapp/service"

	"net/http"

	"github.com/rs/cors"
)

func main() {
	var (
		todoDcl        dcl.TodoDcl               = dcl.New()
		todoService    service.TodoService       = service.New(todoDcl)
		todoController controller.TodoController = controller.New(todoService)
	)

	if err := todoDcl.InitDb(); err != nil {
		log.Fatalln(err)
	}
	defer todoDcl.CloseConnection()
	r := router.Router(todoController)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	}).Handler(r)

	port := ":3000"

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, handler))
}
