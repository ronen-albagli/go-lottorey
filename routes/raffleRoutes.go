package routes

import (
	"lotto/controllers" //add this
	"lotto/service"

	"github.com/gin-gonic/gin"
)

func RaffleRoute(router *gin.Engine) {
	router.Use(service.AuthMiddleWare().MiddlewareFunc())
	router.POST("/raffle", controllers.CreateRaffle())
	router.GET("/raffle/:raffleId", controllers.GetRaffle())  //add this
	router.PUT("/raffle/:raffleId", controllers.EditRaffle()) //add this
	router.GET("/raffle/random-numbers", controllers.GenerateRandomNumbers())
}
