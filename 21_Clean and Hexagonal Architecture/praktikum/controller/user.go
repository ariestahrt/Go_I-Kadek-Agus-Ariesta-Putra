package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	useCase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		userUsecase,
	}
}

func (u *userController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := u.useCase.GetAllUsers()

		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err,
			})
		}

		return c.JSON(200, echo.Map{
			"data": users,
		})
	}
}

func (u *userController) CreateUser() echo.HandlerFunc {
	user := model.User{}
	return func(c echo.Context) error {
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err := u.useCase.CreateUser(user)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err,
			})
		}
		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}

func (u *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		userInput := model.UserInput{}

		if err := c.Bind(&userInput); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "invalid request",
			})
		}

		user, err := u.useCase.Register(userInput)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "register failed",
			})
		}

		return c.JSON(http.StatusCreated, user)
	}
}

func (u *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		userInput := model.UserInput{}

		if err := c.Bind(&userInput); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "invalid request",
			})
		}

		token := u.useCase.Login(userInput)

		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "combination of email and password is wrong",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}

}