package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"belajar-go-echo/util"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user model.User) error
	Register(input model.UserInput) (model.User, error)
	Login(userInput model.UserInput) string
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) GetAllUsers() ([]model.User, error) {
	users, err := u.userRepository.GetAllUsers()

	return users, err
}

func (u *userUsecase) CreateUser(user model.User) error {
	err := u.userRepository.CreateUser(user)

	return err
}

func (u *userUsecase) Register(userInput model.UserInput) (model.User, error) {
	err := userInput.Validate()

	if err != nil {
		return model.User{}, err
	}

	user := u.userRepository.Register(userInput)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUsecase) Login(userInput model.UserInput) string {
	ValidateErr := userInput.Validate()

	if ValidateErr != nil {
		return ""
	}

	user := u.userRepository.Login(userInput)
	WrongPasswordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if WrongPasswordErr != nil {
		return ""
	}

	token := util.CreateToken(user.ID)

	return token
}