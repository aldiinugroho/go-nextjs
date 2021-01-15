package database

import (
	"fmt"
	"rst/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() {
	connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(connection), &gorm.Config{})
	db.Migrator().DropTable("tags", "contents")
	db.Migrator().AutoMigrate(&model.Tag{}, &model.Content{})
	// db.Migrator().DropTable("users")
	// db.Migrator().DropTable("credit_cards")
	var tag = []model.Tag{{TagName: "Golang / Go", TagColor: "#03fcdb"}, {TagName: "JavaScript", TagColor: "#fce303"}, {TagName: "TypeScript", TagColor: "#0384fc"}}
	db.Create(&tag)
	for _, tags := range tag {
		fmt.Println(tags.TagName)
	}
}
