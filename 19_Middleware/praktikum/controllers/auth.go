package controllers

import (
	"net/http"
	"praktikum19/auth"
	"praktikum19/database"
	"praktikum19/models"

	"github.com/labstack/echo"
)


func Login(c echo.Context) error {
	var user *models.User = new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}
	
	var userDB *models.User = new(models.User)
	database.DB.First(&userDB, "email = ?", user.Email)
	
	if userDB.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid username or password",
		})
	}

	token := auth.CreateToken(user.ID)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}