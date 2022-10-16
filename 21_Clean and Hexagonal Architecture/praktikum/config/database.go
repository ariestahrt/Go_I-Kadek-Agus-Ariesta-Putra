package config

import (
	"belajar-go-echo/model"
	"belajar-go-echo/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(util.GetConfig("DB_SQLITE_PATH")), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}