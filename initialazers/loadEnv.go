package initialazers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	var err error = godotenv.Load()
	if err != nil {
		log.Fatal(`error loading .env file`, err.Error())
	}
}
