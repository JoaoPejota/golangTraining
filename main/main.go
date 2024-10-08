package main

import (
	"github.com/JoaoPejota/golangTraining/db"
	"github.com/JoaoPejota/golangTraining/routes"
)

func main(){
	routes.RouterInit()
	db.Connect()
}