package router

import (
	"fmt"
	"net/http"
	"todoapp/controller"

	"github.com/gorilla/mux"
)

func Router(c controller.TodoController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Api is running.")
	})
	router.HandleFunc("/todo", c.GetList).Methods("GET")
	router.HandleFunc("/todo", c.Add).Methods("POST")

	return router
}
