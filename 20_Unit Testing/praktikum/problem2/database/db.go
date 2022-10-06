package database

import (
	"fmt"
	"praktikum19/configs"
	"praktikum19/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfig struct {
	Username string
	Password string
	Port     string
	Host     string
	Name     string
}  

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	config := DBConfig{
	  Username: configs.Env("DB_USERNAME"),
	  Password: configs.Env("DB_PASSWORD"),
	  Port:     configs.Env("DB_PORT"),
	  Host:     configs.Env("DB_HOST"),
	  Name:     configs.Env("DB_NAME"),
	}
   
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	  config.Username,
	  config.Password,
	  config.Host,
	  config.Port,
	  config.Name,
	)
   
	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
	  panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}

