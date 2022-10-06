package controllers

import (
	"net/http"
	"praktikum19/database"
	"praktikum19/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User
	
	if err := database.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "success get all users",
		"data":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	user := models.User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"messages": "input missmatch",
		})
	}

	if err := database.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"message": "failed get user",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "OK",
		"data":   user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "success create new user",
		"data":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"messages": "input missmatch",
		})
	}

	delete := database.DB.Delete(&user, id)

	if delete.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"message": "failed to delete user",
		})
	}

	if delete.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "error",
			"message": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "OK",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	newUser := models.User{}
	c.Bind(&newUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":"error",
			"messages": "input missmatch",
		})
	}

	update := database.DB.Model(&user).Where("id = ?", id).Update(newUser)

	if update.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"message": "failed to update user",
		})
	}

	if update.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "error",
			"message": "user not found",
		})
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "OK",
		"data":   user,
	})
}