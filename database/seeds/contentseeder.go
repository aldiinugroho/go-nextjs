package seeds

import (
	"fmt"
	"os"
	model "rst/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func contentseeder() {
	// connection := "host=localhost user=postgres password=2201 dbname=myweb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		fmt.Println("connecttion error")
	}
	var content = []model.Content{{TagID: 2, ContentTitle: "Fetch data in javascript", ContentContent: "get url insert it into function"},
		{TagID: 2, ContentTitle: "useEffect in react", ContentContent: "call the method use it"},
		{TagID: 1, ContentTitle: "scan white space in golang", ContentContent: "get function ...."}}
	db.Create(&content)
	sqlDB, err := db.DB()
	sqlDB.Close()
}
