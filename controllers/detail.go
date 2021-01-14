package controllers

import "net/http"

func Detail(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client/out/detail.html")
}
