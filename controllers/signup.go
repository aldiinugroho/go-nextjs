package controllers

import (
	"fmt"
	"net/http"
	error "rst/errors"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	error.Check()
	http.ServeFile(w, r, "client/out/signup.html")
}

func Signupdata(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Fprint(w, username+email+password)
}
