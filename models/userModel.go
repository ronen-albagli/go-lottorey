package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	FirstName string             `json:"firstName,omitempty"`
	LastName  string             `json:"lastNameName,omitempty"`
	Password  string             `json:"Password,omitempty"`
	Email     string             `json:"email,omitempty,unique" validate:"required"`
	JoinDate  time.Time          `json:"joinDate,omitempty" validate:"required"`
}
