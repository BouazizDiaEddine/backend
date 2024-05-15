package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	router.POST("/book", controllers.BookCreate)
	router.GET("/books", controllers.BookGetAll)
	router.GET("/book/:id", controllers.BookGet)
	router.PUT("/book/:id", controllers.BookUpdate)
	router.DELETE("/book/:id", controllers.BookDelete)
}
