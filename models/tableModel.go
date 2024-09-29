package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID             primitive.ObjectID `bson:"_id"`
	NumbeeOfGuests *int               `json:"number_of_guests" bson:"number_of_guests" validate:"required"`
	TableNumber    *int               `json:"table_number" bson:"table_number" validate:"required"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
	TableID        string             `json:"table_id" bson:"table_id"`
}
