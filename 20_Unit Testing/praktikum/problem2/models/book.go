package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title         string `json:"title" form:"title"`
	Writer        string `json:"writer" form:"writer"`
	Publisher     string `json:"publisher" form:"publisher"`
	Genre         string `json:"genre" form:"genre"`
	PublishedYear int    `json:"published_year" form:"published_year"`
	TotalPage     int    `json:"total_page" form:"total_page"`
	Price         int    `json:"price" form:"price"`
}