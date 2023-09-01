package database

import (
	"log"

	"github.com/andregoiania/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDados() {
	stringConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	DB.AutoMigrate(&models.Aluno{})
	if err != nil {
		log.Panic("erro ao se conectar ao banco de dados")
	}
}
