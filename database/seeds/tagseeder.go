package seeds

import (
	"fmt"
	"os"
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func tagseeder() {
	// connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		fmt.Println("connecttion error")
	}
	var tag = []model.Tag{{TagName: "Golang", TagColor: "#03fcdb"},
		{TagName: "Javascript", TagColor: "#fce303"}}
	db.Create(&tag)
	sqlDB, err := db.DB()
	sqlDB.Close()
}
