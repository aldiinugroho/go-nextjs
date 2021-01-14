package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func Menupage(w http.ResponseWriter, r *http.Request) {
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
