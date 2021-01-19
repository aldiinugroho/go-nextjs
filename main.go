package main

import (
	"fmt"
	"log"
	"net/http"
	controller "rst/controllers"

	"github.com/joho/godotenv"
)

var port = "3030"

func main() {
	server()
}

func server() {
	// serving dari folder + semua extension atau file yg diperlukan

	// buildHandler := http.FileServer(http.Dir("client/out"))
	// rootfile.Handle("/", buildHandler)
	// http.HandleFunc("/", buildHandler)
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
	// rootfile.Handle("/", http.FileServer(http.Dir("client/out")))
	env()
	controller.Rootindex()

	http.HandleFunc("/inputdata", controller.Inputdata)
	http.HandleFunc("/input", controller.Input)

	http.HandleFunc("/getContent", controller.GetContent)
	http.HandleFunc("/getTag", controller.GetTag)

	http.HandleFunc("/clickTag/", controller.ClickTag)
	http.HandleFunc("/clickTag/getDetail", controller.GetTag_detail)
	http.HandleFunc("/tag", controller.Tag)

	http.HandleFunc("/clickDetail/", controller.ClickDetail)
	http.HandleFunc("/clickDetail/getDetail", controller.GetDetail)
	http.HandleFunc("/detail", controller.Detail)
	// database.Database()
	serve()
}

func serve() {
	fmt.Println("Server run at port : " + port)
	http.ListenAndServe(":"+port, nil)
}

func env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
