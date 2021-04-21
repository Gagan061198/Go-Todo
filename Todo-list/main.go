package main

import (
	"net/http"

	Handlers "TODO-LIST/Handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {

	log.Info("Starting Todolist API server")
	router := mux.NewRouter()

	router.HandleFunc("/todo-completed", Handlers.GetCompletedItems).Methods("GET")

	router.HandleFunc("/todo", Handlers.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", Handlers.UpdateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", Handlers.DeleteItem).Methods("DELETE")

	//handler := cors.New(cors.Options{
	//	AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	//}).Handler(router)

	http.ListenAndServe(":8080", router)
}
