// package config

// import (
// 	"belajar-go-echo/model"
// 	"belajar-go-echo/util"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// func ConnectDB() (*gorm.DB, error) {
// 	return gorm.Open(sqlite.Open(util.GetConfig("DB_SQLITE_PATH")), &gorm.Config{})
// }

// func MigrateDB(db *gorm.DB) error {
// 	return db.AutoMigrate(
// 		model.User{},
// 	)
// }

package config

import (
	"belajar-go-echo/model"
	"belajar-go-echo/util"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := util.GetConfig("DB_USERNAME") + ":" + util.GetConfig("DB_PASSWORD") + "@tcp(" + util.GetConfig("DB_HOST") + ":" + util.GetConfig("DB_PORT") + ")/" + util.GetConfig("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}