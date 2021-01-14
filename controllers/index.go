package controllers

import (
	"net/http"
	"strings"
)

var rootfile = http.NewServeMux()

func Rootindex() {
	rootfile.Handle("/", http.FileServer(http.Dir("client/out")))
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".") {
		rootfile.ServeHTTP(w, r)
	} else {
		if r.URL.Path != "/" {
			http.ServeFile(w, r, "client/out/404.html")
			return
		}
		http.ServeFile(w, r, "client/out/index.html")
	}
}

// func Testjson(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprint(w, "helooo aldi")
// 	// http.ServeFile(w, r, "client/build/index.html")

// 	// sending json
// 	data := UserData{{"armen rais", 20}, {"sri sulistiyowati", 21}, {"eva nur lizar", 22}, {"aldi nugroho", 23}}
// 	js, _ := json.Marshal(data)
// 	// w.Header().Set("Content-Type", "*")
// 	w.Write(js)
// }

// func Clientdata(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprint(w, "oke")
// 	http.ServeFile(w, r, "client/out/clientdata.html")
// }

// func Tescookie(w http.ResponseWriter, r *http.Request) {
// 	// recorder := httptest.NewRecorder()
// 	// request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
// 	getcookie, _ := r.Cookie("userdata")
// 	fmt.Println(getcookie.Value)
// 	// datatosend := []byte(getcookie.Value)
// 	// w.Write(datatosend)
// 	fmt.Fprint(w, getcookie.Value)
// }

// // UserData for ....
// type UserData []struct {
// 	Name string
// 	Age  int
// }

// func Menupage(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	// fmt.Println("===new data===")
// 	// fmt.Println("username : " + r.FormValue("username"))
// 	// fmt.Println("age : " + r.FormValue("age"))
// 	// fmt.Println("==============")
// 	username := r.FormValue("username")
// 	age := r.FormValue("age")
// 	myage, _ := strconv.Atoi(age)
// 	data := UserData{{username, myage}}
// 	js, _ := json.Marshal(data)
// 	datacokiee := string(js)
// 	// w.Write(js)
// 	// remove " error cookie
// 	fixdata := strings.Replace(datacokiee, "\"", "", -1)
// 	cookie := http.Cookie{Name: "userdata", Value: fixdata}
// 	// cookie start
// 	// fmt.Println(cookie)
// 	// recorder := httptest.NewRecorder()
// 	http.SetCookie(w, &cookie)
// 	// request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
// 	// cookies, _ := request.Cookie("userdata")
// 	// fmt.Println(cookies)
// 	// redirect page
// 	// http.Redirect(w, r, "/", http.StatusSeeOther)
// 	// ==========
// 	// recorder := httptest.NewRecorder()
// 	// http.SetCookie(recorder, &http.Cookie{Name: "test", Value: fixdata})
// 	// request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
// 	// getcookie, _ := request.Cookie("test")
// 	// fmt.Println(getcookie.Value)
// 	// datatosend := []byte(getcookie.Value)
// 	// w.Write(datatosend)
// 	http.Redirect(w, r, "/coo", http.StatusSeeOther)
// }
