package main

import (
	"go-todo/routes"
	"log"
	"net/http"
)

func main() {

	err := http.ListenAndServe(":8080", routes.Init())

	if err != nil {

		log.Fatal(err)
	}
}
