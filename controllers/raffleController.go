package controllers

import (
	"context"
	"fmt"
	"lotto/configs"
	"lotto/models"
	"lotto/responses"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var raffleCollection *mongo.Collection = configs.GetCollection(configs.DB, "raffle")
var raffleValidator = validator.New()

func CreateRaffle() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var raffle models.Raffle
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&raffle); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&raffle); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		fmt.Println(raffle.UserEmail)

		newRaffle := models.Raffle{
			Id:         primitive.NewObjectID(),
			UserEmail:  raffle.UserEmail,
			Date:       raffle.Date,
			RaffleDate: raffle.RaffleDate,
			Num1:       raffle.Num1,
			Num2:       raffle.Num2,
			Num3:       raffle.Num3,
			Num4:       raffle.Num4,
			Num5:       raffle.Num5,
			Num6:       raffle.Num6,
			StrongNum:  raffle.StrongNum,
		}

		result, err := raffleCollection.InsertOne(ctx, newRaffle)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetRaffle() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		raffleId := c.Param("raffleId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(raffleId)

		err := raffleCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func EditRaffle() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		raffleId := c.Param("raffleId")
		var raffle models.Raffle
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(raffleId)

		//validate the request body
		if err := c.BindJSON(&raffleId); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := raffleValidator.Struct(&raffle); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"userEmail": raffle.UserEmail, "date": raffle.Date, "raffleDate": raffle.RaffleDate}
		result, err := raffleCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := raffleCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
	}
}

func GenerateRandomNumbers() gin.HandlerFunc {
	return func(c *gin.Context) {
		numbersMap := make(map[int]bool)
		resultsMap := make(map[string][]int)
		strong := make([]int, 0)

		for i := 0; i < 6; i++ {
			number := RandomNumber(1, 37, numbersMap)

			numbersMap[number] = true
		}

		resultsMap["regular"] = ConvertMapToKeysSlice(numbersMap)

		number := RandomNumber(1, 7, numbersMap)

		resultsMap["strong"] = append(strong, number)

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"numbers": resultsMap}})
	}
}

func RandomNumber(min int, max int, numbers map[int]bool) int {
	rand.Seed(time.Now().UnixNano())

	num := (rand.Intn(max-min+1) + min)

	if numbers[num] {
		return RandomNumber(min, max, numbers)
	}

	return num
}

func ConvertMapToKeysSlice(input map[int]bool) []int {
	keys := make([]int, len(input))

	i := 0
	for k := range input {
		keys[i] = k
		i++
	}

	return keys
}
