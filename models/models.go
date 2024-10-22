package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Desc  string  `json:"desc"`
}
