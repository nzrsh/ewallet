package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LoginPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/loginpage.html")
}

func HomePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/homepage.html")
}
func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if FindUserFromUsers(username, password) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		log.Printf("User %s login successfully", username)
	} else {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		log.Printf("User %s failed to login", username)
	}
}
