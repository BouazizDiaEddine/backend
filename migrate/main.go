package main

import (
	"backend/initialazers"
	"backend/models"
)

func init() {
	initialazers.LoadEnvVariables()
	initialazers.ConnectToDb()
}

func main() {
	initialazers.DB.AutoMigrate(&models.Book{}, &models.Url{})
}
