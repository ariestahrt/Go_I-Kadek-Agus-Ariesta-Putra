package routes

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"
	"belajar-go-echo/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userService)

	authenticatedRoutes := e.Group("")
	
	e.POST("/register", userController.Register())
	e.POST("/login", userController.Login())

	authenticatedRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(util.GetConfig("SECRET_JWT_KEY")),
	}))
	
	authenticatedRoutes.GET("/users", userController.GetAllUsers())
	authenticatedRoutes.POST("/users", userController.CreateUser())
}