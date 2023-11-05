package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ronanzindev/echo-fly.io/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectPostgresDB() {
	var stringDeConexao string
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "dev")
	}
	if os.Getenv("ENV") == "prod" {
		stringDeConexao = os.Getenv("CONECTION_STRING")
	}

	if os.Getenv("ENV") == "dev" {
		stringDeConexao = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			"localhost",
			"root",
			"root",
			"enzo-acad",
			"5432",
			"disable",
			"America/Sao_Paulo",
		)
	}
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	fmt.Println("database connected successfully !")

	// Migrate the schema
	DB.AutoMigrate(models.User{})

}
