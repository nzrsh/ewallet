package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var DB *sql.DB

func main() {

	port := flag.Int("port", 8080, "Port for the server to listen on")
	flag.Parse()
	var addr string
	if *port > 0 && *port < 65535 {
		addr = fmt.Sprintf("localhost:%d", *port)
	} else {
		log.Fatalln("Invalid port")
	}

	err := OpenDB()
	if err != nil {
		log.Fatalln(err)
	}

	router := httprouter.New()

	SetupRoutes(router)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("Hosting on:", addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
