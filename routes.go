package main

import "github.com/julienschmidt/httprouter"

func SetupRouter(r *httprouter.Router) {
	r.GET("/", LoginPage) //
	r.GET("/home", HomePage)
	r.POST("/auth", Auth)
}
