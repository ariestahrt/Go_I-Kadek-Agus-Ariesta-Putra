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

func TestCreateBooks_Success(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Body				models.Book
		ExpectedStatus      int
		ExpectBodyStatus string
	}{
		{
			Name:                "Success:: Create book",
			Path:                "/books",
			Method:				http.MethodPost,
			Body:				models.Book{
				Title: "Test Book",
				Writer: "Ariesta",
				Publisher: "PT Cinta Mencari Cinta Sejati",
				Genre: "Komedi",
				PublishedYear: 2022,
				TotalPage: 22,
				Price: 10000,
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

		err := CreateBookController(c)

		assert.NoError(t, err)
		assert.Equal(t, testCase.ExpectedStatus, rec.Code)

		body := rec.Body.String()
		body_json := models.Response{}
		json.Unmarshal([]byte(body), &body_json)
		assert.True(t, body_json.Status == testCase.ExpectBodyStatus)
	}
}

func TestCreateBooks_Failed(t *testing.T) {
	type BookMalformed struct {
		Name     string `json:"name" form:"name"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Title         bool `json:"title" form:"title"`
	}
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Body				BookMalformed
	}{
		{
			Name:                "Failed:: Create book",
			Path:                "/books",
			Method:				http.MethodPost,
			Body:				BookMalformed{
				Name: "Test Book",
				Email: "test@book.com",
				Title: true,
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

		err := CreateBookController(c)
		assert.Error(t, err)
	}
}

func TestGetAllBooks(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		ExpectedStatus      int
		ExpectBodyStatus 	string
	}{
		{
			Name:                "Success:: get all books",
			Path:                "/books",
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
	
			if assert.NoError(t, GetBooksController(c)) {
				assert.Equal(t, testCase.ExpectedStatus, rec.Code)
				body := rec.Body.String()
				body_json := models.Response{}
				json.Unmarshal([]byte(body), &body_json)

				assert.True(t, body_json.Status == testCase.ExpectBodyStatus)
			}
		}
	}
}

func TestGetABooks(t *testing.T) {
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
			Name:                "Success:: Get books with id 1",
			Path:                "/books/:id",
			Method:				http.MethodGet,
			Params:				map[string]string{"id":"1"},
			ExpectedStatus:      http.StatusOK,
			ExpectBodyStatus: "success",
		},
		{
			Name:                "Failed:: Get books with malformed id",
			Path:                "/books/:id",
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

		err := GetBookController(c)
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

func TestUpdateBook(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Body				models.Book
		Params				map[string]string
		HasError			bool
		ExpectedStatus      int
		ExpectBodyStatus string
	}{
		{
			Name:                "Success:: Update book",
			Path:                "/books/:id",
			Method:				http.MethodPost,
			Body:				models.Book{
				Title: "UPDATED Book",
				Writer: "ARIES",
				Publisher: "PT GANTENG",
				Genre: "Komedi",
				PublishedYear: 2022,
				TotalPage: 22,
				Price: 12000,
			},
			Params:				map[string]string{"id":"1"},
			ExpectedStatus:      http.StatusOK,
			ExpectBodyStatus: "success",
		},
		{
			Name:                "Failed:: Update book",
			Path:                "/books/:id",
			Method:				http.MethodPost,
			Body:				models.Book{
				Title: "UPDATED Book 2",
				Writer: "ARIES 2",
				Publisher: "PT GANTENG 2",
				Genre: "Komedi 2",
				PublishedYear: 2022,
				TotalPage: 22,
				Price: 12000,
			},
			Params:				map[string]string{"id":"-1"},
			ExpectedStatus:      http.StatusBadRequest,
			ExpectBodyStatus: 	"error",
		},
		{
			Name:                "Failed:: Update book with malformed id",
			Path:                "/books/:id",
			Method:				http.MethodPost,
			Body:				models.Book{
				Title: "UPDATED Book 3",
				Writer: "ARIES 3",
				Publisher: "PT GANTENG 3",
				Genre: "Komedi 3",
				PublishedYear: 2022,
				TotalPage: 22,
				Price: 12000,
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

		err := UpdateBookController(c)

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

func TestDeleteBook(t *testing.T) {
	var testCases = []struct {
		Name                string
		Path                string
		Method				string
		Body				models.Book
		Params				map[string]string
		HasError			bool
		ExpectedStatus      int
		ExpectBodyStatus string
	}{
		{
			Name:				"Success:: Delete book",
			Path:				"/books/:id",
			Method:				http.MethodDelete,
			Params:				map[string]string{"id":"1"},
			ExpectedStatus:      http.StatusOK,
			ExpectBodyStatus: "success",
		},
		{
			Name:                "Failed:: Delete book",
			Path:                "/books/:id",
			Method:				http.MethodDelete,
			Params:				map[string]string{"id":"-1"},
			ExpectedStatus:      http.StatusBadRequest,
			ExpectBodyStatus: 	"error",
		},
		{
			Name:                "Failed:: Delete book with malformed id",
			Path:                "/books/:id",
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

		err := DeleteBookController(c)

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