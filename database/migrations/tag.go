package migrations

import (
	"fmt"
	"os"
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Tag() {
	// connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		fmt.Println("connecttion error")
	}
	db.Migrator().AutoMigrate(&model.Tag{})
	sqlDB, err := db.DB()
	sqlDB.Close()
}

func DropTag() {
	// connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		fmt.Println("connecttion error")
	}
	db.Migrator().DropTable("tags")
	sqlDB, err := db.DB()
	sqlDB.Close()
}
