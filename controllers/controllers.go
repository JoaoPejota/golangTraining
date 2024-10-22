package controllers

import (
	"net/http"

	"github.com/JoaoPejota/golangTraining/database"
	"github.com/JoaoPejota/golangTraining/models"
	"github.com/gin-gonic/gin"
)

func CriarGame(c *gin.Context) {
	var games models.Game
	if err := c.ShouldBindJSON(&games); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&games)
	c.JSON(http.StatusOK, games)
}

func GetGame(c *gin.Context) {
	var games []models.Game
	database.DB.Find(&games)
	c.JSON(http.StatusOK, games)
}

func GetGamesById(c *gin.Context) {
	var game models.Game
	id := c.Params.ByName("id")
	database.DB.First(&game, id)
	c.JSON(http.StatusOK, game)
}

func DeleteGame(c *gin.Context) {
	var game models.Game
	id := c.Params.ByName("id")
	database.DB.Delete(&game, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso!"})
}

func AtualizarGame(c *gin.Context) {
	var game models.Game
	id := c.Params.ByName("id")
	database.DB.First(&game, id)

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"NÃO DEU CERTO": "A"})
		return
	}
	database.DB.Model(&game).UpdateColumns(game)
	c.JSON(http.StatusOK, game)
}

func GetGameByDesc(c *gin.Context) {
	var game models.Game

	desc := c.Param("desc")
	database.DB.Where(&models.Game{
		Desc: desc}).First(&game)

	c.JSON(http.StatusOK, game)
}
