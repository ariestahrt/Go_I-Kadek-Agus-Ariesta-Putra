package usecase_test

import (
	"belajar-go-echo/model"
	_userMock "belajar-go-echo/repository/mocks"
	usecase "belajar-go-echo/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository _userMock.UserRepository
	userUseCase    usecase.UserUsecase
)

func TestMain(m *testing.M) {
	userUseCase = usecase.NewUserUsecase(&userRepository)

	m.Run()
}

func TestRegister(t *testing.T) {
	type testCases struct {
		name    	string
		event    	string
		input	 	model.UserInput
		assert  	func (t *testing.T, result model.User)
		expected 	model.User
	}
	
	testCase := []testCases{
		{
			name: "Register | Valid",
			event: "Register",
			input: model.UserInput{
				Email: "susanto@gmail.com",
				Password: "123456",
			},
			assert: func (t *testing.T, result model.User) {
				assert.NotNil(t, result)
			},
			expected: model.User{
				Email: "susanto@gmail.com",
				Password: "123456",
			},
		},
		{
			name: "Register | Invalid",
			event: "Register",
			input: model.UserInput{
				Email: "susanto@gmail.com",
				Password: "",
			},
			assert: func (t *testing.T, result model.User) {
				assert.NotNil(t, result)
			},
		},
		{
			name: "Register | Invalid email",
			event: "Register",
			input: model.UserInput{
				Email: "susantogmail.com",
				Password: "123456",
			},
			assert: func (t *testing.T, result model.User) {
				assert.NotNil(t, result)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			User := model.User{}

			userRepository.On(tc.event, tc.input).Return(User, err).Once()
			result, _ := userUseCase.Register(tc.input)

			tc.assert(t, result)
		})
	}
}

func TestLogin(t *testing.T) {
	type testCases struct {
		name     string
		event    string
		input	 model.UserInput
		assert   func (t *testing.T, result string)
	}

	testCase := []testCases{
		{
			name: "Login | Valid",
			event: "Login",
			input: model.UserInput{
				Email: "susanto@gmail.com",
				Password: "123456",
			},
			assert: func (t *testing.T, result string) {
				assert.NotNil(t, result)
			},
		},
		{
			name: "Login | Invalid Password",
			event: "Login",
			input: model.UserInput{
				Email: "susanto@gmail.com",
				Password: "asd123",
			},
			assert: func (t *testing.T, result string) {
				assert.Empty(t, result)
			},
		},
		{
			name: "Login | Invalid Email",
			event: "Login",
			input: model.UserInput{
				Email: "",
				Password: "123456",
			},
			assert: func (t *testing.T, result string) {
				assert.Empty(t, result)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			User := model.User{}

			userRepository.On(tc.event, tc.input).Return(User, err).Once()
			result := userUseCase.Login(tc.input)

			tc.assert(t, result)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	type testCases struct {
		name     string
		event    string
		assert   func (t *testing.T, result []model.User)
	}

	testCase := []testCases{
		{
			name: "GetAllUsers | Valid",
			event: "GetAllUsers",
			assert: func (t *testing.T, result []model.User) {
				assert.NotNil(t, result)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			Users := []model.User{}

			userRepository.On(tc.event).Return(Users, err).Once()
			result, _ := userUseCase.GetAllUsers()

			tc.assert(t, result)
		})
	}
}

func TestCreateUser(t *testing.T) {
	type testCases struct {
		name     string
		event    string
		input	 model.User
		assert   func (t *testing.T, result error)
	}

	testCase := []testCases{
		{
			name: "CreateUser | Valid",
			event: "CreateUser",
			input: model.User{
				Email: "susanto@gmail.com",
				Password: "123456",
			},
			assert: func (t *testing.T, result error) {
				assert.Nil(t, result)
			},
		},
		{
			name: "CreateUser | Invalid",
			event: "CreateUser",
			input: model.User{
				Email: "susantogmail.com",
				Password: "123456",
			},
			assert: func (t *testing.T, result error) {
				assert.Empty(t, result)
			},
		},
	}
	
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			var err error

			userRepository.On(tc.event, tc.input).Return(err).Once()
			result := userUseCase.CreateUser(tc.input)

			tc.assert(t, result)
		})
	}
}
