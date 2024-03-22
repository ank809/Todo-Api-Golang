package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	IsCompleted bool               `json:"iscompleted"`
	CreatedAt   time.Time          `json:"createdat"`
	UpdatedAt   time.Time          `json:"updatedat"`
}
