package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	param "rst/errors"
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Tag(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client/out/tag.html")
}

func GetTag_detail(w http.ResponseWriter, r *http.Request) {
	var contents = []model.Content{}
	getvalue, _ := r.Cookie("contentvalue")
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		fmt.Println("connecttion error")
	}
	db.Preload(clause.Associations).Where("tag_id = ?", getvalue.Value).Find(&contents)
	js, _ := json.Marshal(contents)
	sqlDB, err := db.DB()
	sqlDB.Close()
	w.Write(js)
}

func ClickTag(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	x := param.Parameter_tag(path)
	if x == "0" || x == "" {
		http.ServeFile(w, r, "client/out/404.html")
		return
	}
	cookie := http.Cookie{Name: "contentvalue", Value: x}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/tag", http.StatusSeeOther)
}
