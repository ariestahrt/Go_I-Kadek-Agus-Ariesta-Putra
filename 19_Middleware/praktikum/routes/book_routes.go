package routes

import (
	"praktikum19/configs"
	"praktikum19/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func BookRoute(e *echo.Echo) {
	authenticatedRoutes := e.Group("")

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)

	authenticatedRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(configs.Env("SECRET_KEY")),
	}))

	authenticatedRoutes.POST("/books", controllers.CreateBookController)
	authenticatedRoutes.DELETE("/books/:id", controllers.DeleteBookController)
	authenticatedRoutes.PUT("/books/:id", controllers.UpdateBookController)
}