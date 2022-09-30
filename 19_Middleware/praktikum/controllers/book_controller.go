package controllers

import (
	"net/http"
	"praktikum19/database"
	"praktikum19/models"
	"strconv"

	"github.com/labstack/echo"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book
	
	if err := database.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "success get all books",
		"data":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	book := models.Book{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"messages": "input missmatch",
		})
	}

	if err := database.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"message": "failed get book",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "OK",
		"data":   book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := database.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "success create new book",
		"data":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	var book models.Book
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"messages": "input missmatch",
		})
	}
	
	if err := database.DB.Delete(&book, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"message": "failed to delete book",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "OK",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	var book models.Book
	id, err := strconv.Atoi(c.Param("id"))
	newBook := models.Book{}
	c.Bind(&newBook)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":"error",
			"messages": "input missmatch",
		})
	}

	if err := database.DB.Model(&book).Where("id = ?", id).Update(newBook).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"message": "failed to update book",
		})
	}
	

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"message": "OK",
		"data":   book,
	})
}