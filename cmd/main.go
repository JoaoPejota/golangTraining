package cmd

type game struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Type string `json:"type"`
}

var games = []game{
    {ID: 1, Title: "Valorant", Type: "Fantasy/Action", Price: 20.99},
    {ID: 2, Title: "Counter Strike", Type: "War/Action", Price: 79.99},
    {ID: 3, Title: "Elden Ring", Type: "Fantasy/RPG", Price: 199.99},
}

