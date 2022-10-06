package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"praktikum19/database"
	"praktikum19/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUsers_Success(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Body				models.User
		ExpectedStatus      int
		ExpectBodyStatus string
	}{
		{
			Name:                "Success:: Create user",
			Path:                "/users",
			Method:				http.MethodPost,
			Body:				models.User{
				Name: "Test User",
				Email: "test@user.com",
				Password: "Password123",
			},
			ExpectedStatus:      http.StatusOK,
			ExpectBodyStatus: "success",
		},
	}

	database.InitDB()
	e := echo.New()

	for _, testCase := range(testCases){
		t.Logf("TEST => %s", testCase.Name)
		req_body, _ := json.Marshal(testCase.Body)

		req := httptest.NewRequest(testCase.Method, "/", bytes.NewBuffer(req_body))
		rec := httptest.NewRecorder()

		req.Header.Add("Content-Type", "application/json; charset=UTF-8")

		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)

		err := CreateUserController(c)

		assert.NoError(t, err)
		assert.Equal(t, testCase.ExpectedStatus, rec.Code)

		body := rec.Body.String()
		body_json := models.Response{}
		json.Unmarshal([]byte(body), &body_json)
		assert.True(t, body_json.Status == testCase.ExpectBodyStatus)
	}
}

func TestCreateUsers_Failed(t *testing.T) {
	type UserMalformed struct {
		Name     string `json:"name" form:"name"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password bool `json:"password" form:"password" validate:"required"`
	}
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Body				UserMalformed
	}{
		{
			Name:                "Failed:: Create user",
			Path:                "/users",
			Method:				http.MethodPost,
			Body:				UserMalformed{
				Name: "Test User",
				Email: "test@user.com",
				Password: true,
			},
		},
	}

	database.InitDB()
	e := echo.New()

	for _, testCase := range(testCases){
		t.Logf("TEST => %s", testCase.Name)
		req_body, _ := json.Marshal(testCase.Body)

		req := httptest.NewRequest(testCase.Method, "/", bytes.NewBuffer(req_body))
		rec := httptest.NewRecorder()

		req.Header.Add("Content-Type", "application/json; charset=UTF-8")

		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)

		err := CreateUserController(c)
		assert.Error(t, err)
	}
}

func TestGetAllUsers(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		ExpectedStatus      int
		ExpectBodyStatus 	string
	}{
		{
			Name:                "Success:: get all users",
			Path:                "/users",
			Method:				http.MethodGet,
			ExpectedStatus:      http.StatusOK,
			ExpectBodyStatus: 	"success",
		},
	}

	database.InitDB()
	e := echo.New()

	for _, testCase := range(testCases){
		t.Logf("TEST => %s", testCase.Name)
		req := httptest.NewRequest(testCase.Method, testCase.Path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
	
		for _, testCase := range testCases {
			c.SetPath(testCase.Path)
	
			if assert.NoError(t, GetUsersController(c)) {
				assert.Equal(t, testCase.ExpectedStatus, rec.Code)
				body := rec.Body.String()
				body_json := models.Response{}
				json.Unmarshal([]byte(body), &body_json)

				assert.True(t, body_json.Status == testCase.ExpectBodyStatus)
			}
		}
	}
}

func TestGetAUsers(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Params				map[string]string
		ExpectedStatus      int
		HasError			bool
		ExpectBodyStatus string
	}{
		{
			Name:                "Success:: Get users with id 1",
			Path:                "/users/:id",
			Method:				http.MethodGet,
			Params:				map[string]string{"id":"1"},
			ExpectedStatus:      http.StatusOK,
			ExpectBodyStatus: "success",
		},
		{
			Name:                "Failed:: Get users with malformed id",
			Path:                "/users/:id",
			Method:				http.MethodGet,
			Params:				map[string]string{"id":"abc"},
			ExpectedStatus:      http.StatusInternalServerError,
			HasError: 			false,
			ExpectBodyStatus: "error",
		},
	}

	database.InitDB()
	e := echo.New()

	for _, testCase := range(testCases){
		t.Logf("TEST => %s", testCase.Name)
		req := httptest.NewRequest(testCase.Method, testCase.Path, nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)
		for key,val := range(testCase.Params){
			c.SetParamNames(key)
			c.SetParamValues(val)
		}

		err := GetUserController(c)
		if testCase.HasError {
			assert.Error(t, err)
			continue	
		}

		assert.NoError(t, err)
		assert.Equal(t, testCase.ExpectedStatus, rec.Code)

		body := rec.Body.String()
		body_json := models.Response{}
		json.Unmarshal([]byte(body), &body_json)

		assert.True(t, body_json.Status == testCase.ExpectBodyStatus)
	}
}

func TestUpdateUser(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Body				models.User
		Params				map[string]string
		HasError			bool
		ExpectedStatus      int
		ExpectBodyStatus string
	}{
		{
			Name:                "Success:: Update user",
			Path:                "/users/:id",
			Method:				http.MethodPost,
			Body:				models.User{
				Name: "Updated User",
				Email: "test@user.com",
				Password: "Password123",
			},
			Params:				map[string]string{"id":"1"},
			ExpectedStatus:      http.StatusOK,
			ExpectBodyStatus: "success",
		},
		{
			Name:                "Failed:: Update user",
			Path:                "/users/:id",
			Method:				http.MethodPost,
			Body:				models.User{
				Name: "GAGAL UPDATE",
				Email: "test@user.com",
				Password: "Password123",
			},
			Params:				map[string]string{"id":"-1"},
			ExpectedStatus:      http.StatusBadRequest,
			ExpectBodyStatus: 	"error",
		},
		{
			Name:                "Failed:: Update user with malformed id",
			Path:                "/users/:id",
			Method:				http.MethodPost,
			Body:				models.User{
				Name: "GAGAL UPDATE",
				Email: "test@user.com",
				Password: "Password123",
			},
			Params:				map[string]string{"id":"asd"},
			ExpectedStatus:      http.StatusInternalServerError,
			ExpectBodyStatus: 	"error",
		},
	}
	
	database.InitDB()
	e := echo.New()

	for _, testCase := range(testCases){
		t.Logf("TEST => %s", testCase.Name)
		req_body, _ := json.Marshal(testCase.Body)

		req := httptest.NewRequest(testCase.Method, "/", bytes.NewBuffer(req_body))
		rec := httptest.NewRecorder()

		req.Header.Add("Content-Type", "application/json; charset=UTF-8")

		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)
		for key,val := range(testCase.Params){
			c.SetParamNames(key)
			c.SetParamValues(val)
		}

		err := UpdateUserController(c)

		if testCase.HasError {
			assert.Error(t, err)
			continue	
		}

		assert.NoError(t, err)
		assert.Equal(t, testCase.ExpectedStatus, rec.Code)

		body := rec.Body.String()
		body_json := models.Response{}
		json.Unmarshal([]byte(body), &body_json)

		assert.True(t, body_json.Status == testCase.ExpectBodyStatus)
	}
}


func TestLogin(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Body				models.User
		HasError			bool
		ExpectedStatus      int
		ExpectBodyStatus string
	}{
		{
			Name:                "Success:: Login",
			Path:                "/login",
			Method:				http.MethodPost,
			Body:				models.User{
				Email: "test@user.com",
				Password: "Password123",
			},
			ExpectedStatus:      http.StatusOK,
		},
		{
			Name:                "Failed:: Login",
			Path:                "/login",
			Method:				http.MethodPost,
			Body:				models.User{
				Email: "test@user.com",
				Password: "Password123x",
			},
			ExpectedStatus:      http.StatusUnauthorized,
		},
	}
	
	database.InitDB()
	e := echo.New()

	for _, testCase := range(testCases){
		t.Logf("TEST => %s", testCase.Name)
		req_body, _ := json.Marshal(testCase.Body)

		req := httptest.NewRequest(testCase.Method, "/", bytes.NewBuffer(req_body))
		rec := httptest.NewRecorder()

		req.Header.Add("Content-Type", "application/json; charset=UTF-8")

		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)

		err := Login(c)

		if testCase.HasError {
			assert.Error(t, err)
			continue	
		}

		assert.NoError(t, err)
		assert.Equal(t, testCase.ExpectedStatus, rec.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Body				models.User
		Params				map[string]string
		HasError			bool
		ExpectedStatus      int
		ExpectBodyStatus string
	}{
		{
			Name:				"Success:: Delete user",
			Path:				"/users/:id",
			Method:				http.MethodDelete,
			Params:				map[string]string{"id":"1"},
			ExpectedStatus:      http.StatusOK,
			ExpectBodyStatus: "success",
		},
		{
			Name:                "Failed:: Delete user",
			Path:                "/users/:id",
			Method:				http.MethodDelete,
			Params:				map[string]string{"id":"-1"},
			ExpectedStatus:      http.StatusBadRequest,
			ExpectBodyStatus: 	"error",
		},
		{
			Name:                "Failed:: Delete user with malformed id",
			Path:                "/users/:id",
			Method:				http.MethodPost,
			Params:				map[string]string{"id":"asd"},
			ExpectedStatus:      http.StatusInternalServerError,
			ExpectBodyStatus: 	"error",
		},
	}
	
	database.InitDB()
	e := echo.New()

	for _, testCase := range(testCases){
		t.Logf("TEST => %s", testCase.Name)
		req_body, _ := json.Marshal(testCase.Body)

		req := httptest.NewRequest(testCase.Method, "/", bytes.NewBuffer(req_body))
		rec := httptest.NewRecorder()

		req.Header.Add("Content-Type", "application/json; charset=UTF-8")

		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)
		for key,val := range(testCase.Params){
			c.SetParamNames(key)
			c.SetParamValues(val)
		}

		err := DeleteUserController(c)

		if testCase.HasError {
			assert.Error(t, err)
			continue	
		}

		assert.NoError(t, err)
		assert.Equal(t, testCase.ExpectedStatus, rec.Code)

		body := rec.Body.String()
		body_json := models.Response{}
		json.Unmarshal([]byte(body), &body_json)

		assert.True(t, body_json.Status == testCase.ExpectBodyStatus)
	}
}