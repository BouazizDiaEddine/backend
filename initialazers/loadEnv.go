package initialazers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	var err error = godotenv.Load()
	if err != nil {
		Logger.Print(`error loading .env file`, err.Error())
	}
}
