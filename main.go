package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var port = "3000"

func main() {
	server()
}

func server() {
	// serving dari folder + semua extension atau file yg diperlukan
	// buildHandler := http.FileServer(http.Dir("client/out"))
	// http.Handle("/", http.FileServer(http.Dir("client/out")))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.FileServer(http.Dir("client/out"))
	// 	if r.URL.Path != "/" {
	// 		http.ServeFile(w, r, "client/out/404.html")
	// 		return
	// 	}
	// 	http.ServeFile(w, r, "client/out/index.html")
	// })
	// // serving dari function
	http.HandleFunc("/", index)
	http.HandleFunc("/data", testjson)
	http.HandleFunc("/menu", menupage)
	serve()
}

func index(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("client/out"))
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "client/out/404.html")
		return
	}
	http.ServeFile(w, r, "client/out/index.html")
}

type UserData []struct {
	Name string
}

func testjson(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "helooo aldi")
	// http.ServeFile(w, r, "client/build/index.html")

	// sending json
	data := UserData{{"armen"}, {"sri"}, {"eva"}, {"aldi"}}
	js, _ := json.Marshal(data)
	// w.Header().Set("Content-Type", "*")
	w.Write(js)
}

func menupage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client/out/menu.html")
}

func serve() {
	fmt.Println("Server run at port : " + port)
	http.ListenAndServe(":"+port, nil)
}
