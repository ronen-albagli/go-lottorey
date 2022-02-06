package service

import (
	"context"
	"fmt"
	"lotto/configs"
	"lotto/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewUser struct {
	Email    string
	JoinDate time.Time
}

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func GenreateUser(c *gin.Context) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		return nil, err
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return nil, validationErr
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Email:    user.Email,
		JoinDate: user.JoinDate,
	}

	fmt.Println(user)

	result, err := userCollection.InsertOne(ctx, newUser)

	fmt.Println(result)

	return result, err

}
