package repository

import (
	"belajar-go-echo/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user model.User) error
	Register(input model.UserInput) model.User
	Login(input model.UserInput) model.User
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (u *UserRepositoryImpl) GetAllUsers() ([]model.User, error) {
	users := make([]model.User, 0)

	if err := u.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) CreateUser(user model.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryImpl) Register(input model.UserInput) model.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser model.User = model.User{
		Email:    input.Email,
		Password: string(password),
	}

	var createdUser model.User = model.User{}

	result := u.db.Create(&newUser)

	result.Last(&createdUser)

	return createdUser
}

func (u *UserRepositoryImpl) Login(input model.UserInput) model.User {
	var user model.User = model.User{}

	u.db.First(&user, "email = ?", input.Email)

	if user.ID == 0 {
		return model.User{}
	}

	return user
}