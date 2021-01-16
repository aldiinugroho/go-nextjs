package migrations

import (
	"os"
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Content() {
	// connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	db.Migrator().AutoMigrate(&model.Content{})
}

func DropContent() {
	// connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	db.Migrator().DropTable("contents")
}
