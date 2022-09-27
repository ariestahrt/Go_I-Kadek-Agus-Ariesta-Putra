package database

import (
	"fmt"
	"praktikum18/models"

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
	  Username: "root",
	  Password: "",
	  Port:     "3306",
	  Host:     "localhost",
	  Name:     "alta_2",
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

