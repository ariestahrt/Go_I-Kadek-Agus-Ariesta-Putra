package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error",
		})
	}

	for _, user := range(users) {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success",
				"users":    user,
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user not exist",
		"users": nil,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error",
		})
	}

	for idx, user := range(users) {
		if user.Id == id {
			users = append(users[:idx], users[idx+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success",
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user not exist",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here

	newUser := User{}
	c.Bind(&newUser)

	id, err := strconv.Atoi(c.Param("id"))
	newUser.Id = id
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error",
		})
	}

	for idx, user := range(users) {
		if user.Id == id {
			users[idx] = newUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user not exist",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
