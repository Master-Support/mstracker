package config

import (
	"log"
	"github.com/joho/godotenv"
)

func CarregaVariaveisDeAmbiente() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}