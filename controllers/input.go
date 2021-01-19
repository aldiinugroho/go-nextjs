package controllers

import (
	"fmt"
	"net/http"
	"os"
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Input(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client/out/input.html")
}

func Inputdata(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("message")
	tagname := r.FormValue("tag")
	tagcolor := r.FormValue("tagcolor")
	// open db
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		fmt.Println("connecttion error")
	}
	var tags model.Tag
	db.Where("tag_name LIKE ?", tagname).Find(&tags)
	if tagname == tags.TagName {
		var newcontent = []model.Content{{TagID: int(tags.TagID), ContentTitle: title, ContentContent: content}}
		db.Create(&newcontent)
	} else {
		// input tags baru
		var newcontent = []model.Content{{TagID: int(tags.TagID), ContentTitle: title, ContentContent: content, Tag: model.Tag{TagName: tagname, TagColor: tagcolor}}}
		db.Create(&newcontent)
	}
	// close db
	sqlDB, err := db.DB()
	sqlDB.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}