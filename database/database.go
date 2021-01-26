package oke

import (
	migration "rst/database/migrations"
)

func Database() {
	// db.Migrator().DropTable("tags", "contents")
	// db.Migrator().AutoMigrate(&model.Tag{}, &model.Content{})
	// db.Migrator().DropTable("users")
	// db.Migrator().DropTable("credit_cards")
	// for _, tags := range tag {
	// 	fmt.Println(tags.TagName)
	// }
	// query data
	// var tags []model.Tag
	// db.Find(&tags)
	// db.Where("tag_name = ?", "TypeScript").Find(&tags)
	// fmt.Println(tags)
	// fmt.Println(reflect.ValueOf(db).Kind())
	migration.DropTag()
	migration.DropContent()
	migration.Tag()
	migration.Content()
	// seed.Seeder()
}
