package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email,omitempty,unique" validate:"required"`
	JoinDate time.Time          `json:"joinDate,omitempty" validate:"required"`
}
