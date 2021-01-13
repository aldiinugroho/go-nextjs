package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	http.HandleFunc("/", index)
	http.HandleFunc("/data", testjson)
	http.HandleFunc("/menu", menupage)
	http.HandleFunc("/coo", tescookie)
	http.HandleFunc("/test", clientdata)
	serve()
}

func index(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".") {
		chttp.ServeHTTP(w, r)
	} else {
		if r.URL.Path != "/" {
			http.ServeFile(w, r, "client/out/404.html")
			return
		}
		http.ServeFile(w, r, "client/out/index.html")
	}
}

type UserData []struct {
	Name string
	Age  int
}

func clientdata(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "oke")
	http.ServeFile(w, r, "client/out/clientdata.html")
}

func testjson(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "helooo aldi")
	// http.ServeFile(w, r, "client/build/index.html")

	// sending json
	data := UserData{{"armen", 20}, {"sri", 21}, {"eva", 22}, {"aldi", 23}, {"i love my self very much ty god", 87787666008}}
	js, _ := json.Marshal(data)
	// w.Header().Set("Content-Type", "*")
	w.Write(js)
}

func menupage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// fmt.Println("===new data===")
	// fmt.Println("username : " + r.FormValue("username"))
	// fmt.Println("age : " + r.FormValue("age"))
	// fmt.Println("==============")
	username := r.FormValue("username")
	age := r.FormValue("age")
	myage, _ := strconv.Atoi(age)
	data := UserData{{username, myage}}
	js, _ := json.Marshal(data)
	datacokiee := string(js)
	// w.Write(js)
	// remove " error cookie
	fixdata := strings.Replace(datacokiee, "\"", "", -1)
	cookie := http.Cookie{Name: "userdata", Value: fixdata}
	// cookie start
	// fmt.Println(cookie)
	// recorder := httptest.NewRecorder()
	http.SetCookie(w, &cookie)
	// request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
	// cookies, _ := request.Cookie("userdata")
	// fmt.Println(cookies)
	// redirect page
	// http.Redirect(w, r, "/", http.StatusSeeOther)
	// ==========
	// recorder := httptest.NewRecorder()
	// http.SetCookie(recorder, &http.Cookie{Name: "test", Value: fixdata})
	// request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
	// getcookie, _ := request.Cookie("test")
	// fmt.Println(getcookie.Value)
	// datatosend := []byte(getcookie.Value)
	// w.Write(datatosend)
	http.Redirect(w, r, "/coo", http.StatusSeeOther)
}

func tescookie(w http.ResponseWriter, r *http.Request) {
	// recorder := httptest.NewRecorder()
	// request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
	getcookie, _ := r.Cookie("userdata")
	fmt.Println(getcookie.Value)
	// datatosend := []byte(getcookie.Value)
	// w.Write(datatosend)
	fmt.Fprint(w, getcookie.Value)
}

func serve() {
	fmt.Println("Server run at port : " + port)
	http.ListenAndServe(":"+port, nil)
}
