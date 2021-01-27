package controllers

import (
	"fmt"
	"net/http"
	error "rst/errors"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	error.Check()
	http.ServeFile(w, r, "client/out/signin.html")
}

func Signindata(w http.ResponseWriter, r *http.Request) {
	error.Check()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Fprint(w, username+password)
}
