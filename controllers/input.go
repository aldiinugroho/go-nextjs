package controllers

import (
	"fmt"
	"net/http"
	"os"
	error "rst/errors"
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Input(w http.ResponseWriter, r *http.Request) {
	error.Check()
	http.ServeFile(w, r, "client/out/input.html")
}

func Inputdata(w http.ResponseWriter, r *http.Request) {
	error.Check()
	title := r.FormValue("title")
	content := r.FormValue("message")
	tagname := r.FormValue("tag")
	tagcolor := "#ffffff"
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
