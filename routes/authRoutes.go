package routes

import (
	// "lotto/controllers" //add this
	"lotto/service"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/login", service.AuthMiddleWare().LoginHandler)
}
