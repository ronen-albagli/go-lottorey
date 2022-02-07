package service

import (
	"context"
	"fmt"
	"lotto/configs"
	"lotto/models"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var raffleCollection *mongo.Collection = configs.GetCollection(configs.DB, "raffle")
var raffleValidator = validator.New()

func CreateRaffle(raffle models.Raffle) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		return result, err
	}

	return result, err
}

func GetRaffle(id string) (models.Raffle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var raffle models.Raffle
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	fmt.Println(objId)

	err := raffleCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&raffle)
	if err != nil {
		return raffle, err
	}

	return raffle, nil

}

func GenerateRandomNumbers() map[string][]int {
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

	return resultsMap
}

func UpdateRaffle(id string, c *gin.Context) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var raffle models.Raffle
	defer cancel()
	fmt.Println(id)

	objId, _ := primitive.ObjectIDFromHex(id)

	//validate the request body
	if err := c.BindJSON(&raffle); err != nil {
		return nil, err
	}

	if validationErr := raffleValidator.Struct(&raffle); validationErr != nil {
		return nil, validationErr
	}

	updatedRaffle := bson.M{
		"userEmail":  raffle.UserEmail,
		"date":       raffle.Date,
		"raffleDate": raffle.RaffleDate,
		"num1":       raffle.Num1,
		"num2":       raffle.Num2,
		"num3":       raffle.Num3,
		"num4":       raffle.Num4,
		"num5":       raffle.Num5,
		"num6":       raffle.Num6,
		"strongNum":  raffle.StrongNum,
	}

	result, err := raffleCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": updatedRaffle})
	if err != nil {
		return result, err
	}

	return result, nil
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
