package routes

import (
	"praktikum18/controllers"

	"github.com/labstack/echo"
)

func BookRoute(e *echo.Echo) {
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
}