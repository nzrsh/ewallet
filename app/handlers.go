package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Send(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
}

// start
func GetLast(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func GetBalance(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
