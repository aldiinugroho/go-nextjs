package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	error "rst/errors"
	param "rst/errors"
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Detail(w http.ResponseWriter, r *http.Request) {
	error.Check()
	http.ServeFile(w, r, "client/out/detail.html")
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	error.Check()
	var contents = model.Content{}
	getdetail, _ := r.Cookie("contentvalue")
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		fmt.Println("connecttion error")
	}
	db.Preload(clause.Associations).Where("content_id = ?", getdetail.Value).Order("created_at desc").Find(&contents)
	js, _ := json.Marshal(contents)
	sqlDB, err := db.DB()
	sqlDB.Close()
	w.Write(js)
}

func ClickDetail(w http.ResponseWriter, r *http.Request) {
	error.Check()
	path := r.URL.Path
	x := param.Parameter_detail(path)

	if x == "0" || x == "" {
		http.ServeFile(w, r, "client/out/404.html")
		return
	}
	cookie := http.Cookie{Name: "contentvalue", Value: x}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/detail", http.StatusSeeOther)
}
