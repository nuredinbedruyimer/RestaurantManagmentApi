package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItems struct {
	ID          primitive.ObjectID `bson:"_id"`
	Quantity    *string            `json:"quantity" bson:"quantity" validate:"required,eq=SMALL|eq=MEDIUM|eq=LARGE"`
	UnitPrice   *float64           `json:"unit_price" bson:"unit_price" validate:"required"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	FoodID      *string            `json:"food_id" bson:"food_id" validate:"required"`
	OrderItemID string             `json:"order_item_id" bson:"order_item_id"`
	OrderID     string             `json:"order_id" bson:"order_id" validate:"required"`
}
