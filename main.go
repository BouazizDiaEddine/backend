package main

import (
	"backend/initialazers"
	"backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initialazers.LoadEnvVariables()
	initialazers.ConnectToDb()
}

func main() {
	router := gin.New()
	routes.BookRoutes(router)
	routes.UrlRoutes(router)
	router.Run()
}
