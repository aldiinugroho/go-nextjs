package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var chttp = http.NewServeMux()

func Index(w http.ResponseWriter, r *http.Request) {
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

func Testjson(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "helooo aldi")
	// http.ServeFile(w, r, "client/build/index.html")

	// sending json
	data := UserData{{"armen rais", 20}, {"sri sulistiyowati", 21}, {"eva nur lizar", 22}, {"aldi nugroho", 23}}
	js, _ := json.Marshal(data)
	// w.Header().Set("Content-Type", "*")
	w.Write(js)
}

func Clientdata(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "oke")
	http.ServeFile(w, r, "client/out/clientdata.html")
}

func Tescookie(w http.ResponseWriter, r *http.Request) {
	// recorder := httptest.NewRecorder()
	// request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
	getcookie, _ := r.Cookie("userdata")
	fmt.Println(getcookie.Value)
	// datatosend := []byte(getcookie.Value)
	// w.Write(datatosend)
	fmt.Fprint(w, getcookie.Value)
}

// UserData for ....
type UserData []struct {
	Name string
	Age  int
}
