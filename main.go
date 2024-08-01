package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Users []User

func main() {
	InitDB()

	router := httprouter.New()
	SetupRouter(router)

	log.Println("Starting server...")

	http.ListenAndServe(":8080", router)
}
