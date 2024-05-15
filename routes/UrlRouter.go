package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func UrlRoutes(router *gin.Engine) {
	router.POST("/url", controllers.UrlProcess)
}
