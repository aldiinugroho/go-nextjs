package main

import (
	"fmt"
	"net/http"
	controller "rst/controllers"
)

var port = "3030"

func main() {
	server()
}

var chttp = http.NewServeMux()

func server() {
	// serving dari folder + semua extension atau file yg diperlukan
	// buildHandler := http.FileServer(http.Dir("client/out"))
	// http.Handle("/", buildHandler)
	// http.Handle("/static/", http.StripPrefix("/static/", buildHandler))
	// http.Handle("/", http.FileServer(http.Dir("client/out")))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.URL.Path != "/" {
	// 		http.ServeFile(w, r, "client/out/404.html")
	// 		return
	// 	}
	// 	http.ServeFile(w, r, "client/out/index.html")
	// })
	// serving dari function
	// http.HandleFunc("/", index)
	// http.Handle("/static/", http.StripPrefix("/static/", buildHandler))
	chttp.Handle("/", http.FileServer(http.Dir("client/out")))
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/data", controller.Testjson)
	http.HandleFunc("/menu", controller.Menupage)
	http.HandleFunc("/coo", controller.Tescookie)
	http.HandleFunc("/test", controller.Clientdata)
	serve()
}

func serve() {
	fmt.Println("Server run at port : " + port)
	http.ListenAndServe(":"+port, nil)
}
