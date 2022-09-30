package routes

import (
	"praktikum19/configs"
	"praktikum19/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func UserRoute(e *echo.Echo) {
	authenticatedRoutes := e.Group("")

	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.Login)

	authenticatedRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(configs.Env("SECRET_KEY")),
	}))

	authenticatedRoutes.GET("/users", controllers.GetUsersController)
	authenticatedRoutes.GET("/users/:id", controllers.GetUserController)
	authenticatedRoutes.DELETE("/users/:id", controllers.DeleteUserController)
	authenticatedRoutes.PUT("/users/:id", controllers.UpdateUserController)
}