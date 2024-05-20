package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	router.POST("/books", controllers.BookCreate)
	router.GET("/books", controllers.BookGetAll)
	router.GET("/books/:id", controllers.BookGet)
	router.PUT("/books/:id", controllers.BookUpdate)
	router.DELETE("/books/:id", controllers.BookDelete)
}
