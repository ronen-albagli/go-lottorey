package main

import (
	"lotto/configs" //add this
	middleware "lotto/middlewares"
	"lotto/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.CORS())

	configs.ConnectDB()
	routes.AuthRoutes(router)
	routes.UserRoute(router)
	routes.RaffleRoute(router)
	router.Run("localhost:3001")
}
