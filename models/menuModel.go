package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name" bson:"name" validate:"required"`
	Category  *string            `json:"category" bson:"category" validate:"required"`
	StartDate *time.Time         `json:"start_date" bson:"start_date"`
	EndDate   *time.Time         `json:"end_date"  bson:"end_date"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	MenuID    string             `json:"menu_id" bson:"menu_id"`
}
