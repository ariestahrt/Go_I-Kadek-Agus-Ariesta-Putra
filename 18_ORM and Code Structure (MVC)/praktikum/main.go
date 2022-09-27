package main

import (
	"praktikum18/database"
	"praktikum18/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
  // init database
  database.InitDB()
  // create a new echo instance
  e := echo.New()
  routes.BookRoute(e)
  routes.UserRoute(e)
  
  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}