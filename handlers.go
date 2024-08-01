package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func StartPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println(fmt.Printf("Startpage %s request was received from %s", r.Header, r.RemoteAddr))
	w.Write([]byte("Startpage"))
}
