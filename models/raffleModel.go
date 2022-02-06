package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Raffle struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	UserEmail  string             `json:"userEmail,omitempty" validate:"required" bson:"userEmail" `
	Date       time.Time          `json:"date,omitempty" validate:"required"`
	RaffleDate time.Time          `json:"raffleDate,omitempty" validate:"required" bson:"RaffleDate"`
	Num1       int                `json:"num1,omitempty" validate:"required" bson:"num1"`
	Num2       int                `json:"num2,omitempty" validate:"required" bson:"num2"`
	Num3       int                `json:"num3,omitempty" validate:"required" bson:"num3"`
	Num4       int                `json:"num4,omitempty" validate:"required" bson:"num4"`
	Num5       int                `json:"num5,omitempty" validate:"required" bson:"num5"`
	Num6       int                `json:"num6,omitempty" validate:"required" bson:"num6"`
	StrongNum  int                `json:"strongNum,omitempty" validate:"required" bson:"strongNum"`
}
