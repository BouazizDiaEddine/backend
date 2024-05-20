package initialazers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	var dsn string = os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Print("could not connect to db", err.Error())
	}
}
