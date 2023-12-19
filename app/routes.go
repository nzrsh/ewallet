package main

import "github.com/julienschmidt/httprouter"

func SetupRoutes(r *httprouter.Router) {
	r.POST("/api/send", Send)
	r.GET("/api/transactions", GetLast)
	r.GET("/api/wallet/:address/balance", GetBalance)
}
