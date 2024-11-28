package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

"os.GetEnv("HOST")"
func ConectaComBancoDeDados() {
	stringDeConexao := "host="+os.GetEnv("HOST")+" user="+os.GetEnv("USER")+" password="+os.GetEnv("PASSWORD")+" dbname="+os.GetEnv("DBNAME")+" port="+os.GetEnv("PORT")+" sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
