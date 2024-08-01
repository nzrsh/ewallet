package main

import "github.com/julienschmidt/httprouter"

func SetupRouter(r *httprouter.Router) {
	r.GET("/", StartPage)
}
