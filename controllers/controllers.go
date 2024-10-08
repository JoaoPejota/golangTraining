package controllers

import (
	"net/http"

	"github.com/JoaoPejota/golangTraining/models"
	"github.com/gin-gonic/gin"
)


var games = []models.Game{
    {ID: "1", Title: "Valorant", Type: "Fantasy/Action", Price: 20.99},
    {ID: "2", Title: "Counter Strike", Type: "War/Action", Price: 79.99},
    {ID: "3", Title: "Elden Ring", Type: "Fantasy/RPG", Price: 199.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}


func GetAlbumsById(c *gin.Context){
	id := c.Param("id")

	for _, game := range games{
		if game.ID == id{
			c.IndentedJSON(http.StatusOK, game)
			return
		}
	}
}

func PostGames(c *gin.Context){
	var newGame models.Game

	if err := c.BindJSON(&newGame); err != nil{
		return
	}

	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, games)
}