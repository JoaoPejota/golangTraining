package models


type Game struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Type string `json:"type"`
}