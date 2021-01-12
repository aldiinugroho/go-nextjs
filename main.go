package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port = "3000"

func main() {
	server()
}

type UserData []struct {
	Name string
}

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "helooo aldi")
	// http.ServeFile(w, r, "client/build/index.html")

	// sending json
	data := UserData{{"armen"}, {"sri"}, {"eva"}, {"aldi"}}
	js, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "*")
	w.Write(js)
}

func server() {
	// serving dari folder + semua extension atau file yg diperlukan
	buildHandler := http.FileServer(http.Dir("client/build"))
	http.Handle("/", buildHandler)
	// serving dari function
	http.HandleFunc("/function", index)
	serve()
}

func serve() {
	fmt.Println("Server run at port : " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
