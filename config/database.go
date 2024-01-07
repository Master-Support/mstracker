package config

import (
	"fmt"
	"log"
	"os"
	"example.com/mstracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("user=%s password=%s dbname=%s host=%s",

			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_HOST")),
			
	}), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.StatusObjeto{})

    fmt.Println("Conexão bem-sucedida!")

    return db
}