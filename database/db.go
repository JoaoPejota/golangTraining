package database

import (
	"log"

	"github.com/JoaoPejota/golangTraining/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao tentar conectar ao banco de dados")
	}
	DB.AutoMigrate(&models.Game{})
}
