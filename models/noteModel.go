package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id"`
	Text      *string            `json:"text" bson:"text" validate:"required"`
	Title     *string            `json:"title" bson:"title" validate:"required,min=2,max=100"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	NoteID    string             `json:"note_id" bson:"note_id"`
}
