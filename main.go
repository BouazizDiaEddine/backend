package main

import (
	"backend/initialazers"
	"backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initialazers.InitLogger()
	initialazers.LoadEnvVariables()
	initialazers.ConnectToDb()
	initialazers.InitValidator()
}

func main() {
	router := gin.New()
	config := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
	}

	// Apply the CORS middleware to the router
	router.Use(cors.New(config))
	routes.BookRoutes(router)
	routes.UrlRoutes(router)
	router.Run()
}
