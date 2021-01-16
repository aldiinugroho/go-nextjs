package migrations

import (
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Tag() {
	connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(connection), &gorm.Config{})
	db.Migrator().AutoMigrate(&model.Tag{})
}

func DropTag() {
	connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(connection), &gorm.Config{})
	db.Migrator().DropTable("tags")
}
