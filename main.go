package main

import (
	"github.com/JoaoPejota/golangTraining/database"
	"github.com/JoaoPejota/golangTraining/routes"
)

func main() {
	database.Connection()
	routes.RouterInit()
}
