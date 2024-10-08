package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/JoaoPejota/golangTraining/controllers"

)

func RouterInit() {
	Router := gin.Default()
	Router.GET("/games", controllers.GetAlbums)
	Router.GET("/games/:id", controllers.GetAlbumsById)
	Router.POST("/games", controllers.PostGames)
	Router.Run("localhost:8080")
}
