package seeds

import (
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func tagseeder() {
	connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(connection), &gorm.Config{})
	var tag = []model.Tag{{TagName: "Golang / Go", TagColor: "#03fcdb"},
		{TagName: "JavaScript", TagColor: "#fce303"},
		{TagName: "TypeScript", TagColor: "#0384fc"}}
	db.Create(&tag)
}
