package controllers

import (
	"fmt"
	"lotto/configs"
	"lotto/models"
	"lotto/responses"
	"lotto/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var raffleCollection *mongo.Collection = configs.GetCollection(configs.DB, "raffle")
var raffleValidator = validator.New()

func CreateRaffle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var raffle models.Raffle

		//validate the request body
		if err := c.BindJSON(&raffle); err != nil {
			c.JSON(http.StatusBadRequest, responses.RaffleResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		result, err := service.CreateRaffle(raffle)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RaffleResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.RaffleResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetRaffle() gin.HandlerFunc {
	return func(c *gin.Context) {
		raffleId := c.Param("raffleId")
		fmt.Println(raffleId)
		result, err := service.GetRaffle(raffleId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RaffleResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.RaffleResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func EditRaffle() gin.HandlerFunc {
	return func(c *gin.Context) {
		raffleId := c.Param("raffleId")

		result, err := service.UpdateRaffle(raffleId, c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RaffleResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.RaffleResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GenerateRandomNumbers() gin.HandlerFunc {
	return func(c *gin.Context) {

		numbers := service.GenerateRandomNumbers()

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"numbers": numbers}})
	}
}
