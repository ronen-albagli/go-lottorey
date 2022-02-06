package routes

import (
	"lotto/controllers" //add this

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())  //add this
	router.PUT("/user/:userId", controllers.EditAUser()) //add this

}
