package routes

import (
	todo "go-todo/Handlers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", todo.Show).Methods("GET")
	route.HandleFunc("/add", todo.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", todo.Delete).Methods("DELETE")
	route.HandleFunc("/complete/{id}", todo.Complete).Methods("GET")

	return route
}
