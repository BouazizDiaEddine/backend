package main

import (
	_ "backend/docs"
	"backend/initialazers"
	"backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initialazers.InitLogger()
	initialazers.LoadEnvVariables()
	initialazers.ConnectToDb()
	initialazers.InitValidator()
}

// @title byfood Library + url treatment
// @version 1.0
// @description crud book library + some URL treatment operation
// @host localhost:8080
func main() {
	router := gin.New()
	config := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
	}
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.New(config))
	routes.BookRoutes(router)
	routes.UrlRoutes(router)
	router.Run()
}
