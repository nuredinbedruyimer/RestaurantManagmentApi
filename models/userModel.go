package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastName     *string            `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Email        *string            `json:"email" bson:"email" validate:"required, email,min=4"`
	Password     *string            `json:"password" bson:"password"`
	Avatar       *string            `json:"avatar" bson:"avatar"`
	PhoneNumber  *string            `json:"phone_number" bson:"phone_number" validate:"required"`
	Token        *string            `json:"token" bson:"token"`
	RefreshToken *string            `json:"refresh_token" bson:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	UserID       string             `json:"user_id" bson:"user_id" validate:"required"`
}
