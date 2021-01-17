package migrations

import (
	"fmt"
	"os"
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Content() {
	// connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		fmt.Println("connecttion error")
	}
	db.Migrator().AutoMigrate(&model.Content{})
	sqlDB, err := db.DB()
	sqlDB.Close()
}

func DropContent() {
	// connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		fmt.Println("connecttion error")
	}
	db.Migrator().DropTable("contents")
	sqlDB, err := db.DB()
	sqlDB.Close()
}
