package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name" validate:"required"`
	Email        string             `bson:"email" json:"email" `
	Password     string             `bson:"password" json:"password" validate:"required"`
	Token        string             `bson:"token" json:"token"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}
