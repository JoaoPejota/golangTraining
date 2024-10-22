package routes

import (
	"github.com/JoaoPejota/golangTraining/controllers"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	Router := gin.Default()
	Router.GET("/games", controllers.GetGame)
	Router.GET("/games/:id", controllers.GetGamesById)
	Router.POST("/games", controllers.CriarGame)
	Router.DELETE("/games/:id", controllers.DeleteGame)
	Router.PATCH("/games/:id", controllers.AtualizarGame)
	Router.GET("/games/desc/:desc", controllers.GetGameByDesc)
	Router.Run("localhost:8080")
}
