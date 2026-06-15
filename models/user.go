package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName      string             `bson:"full_name" json:"full_name"`
	Email         string             `bson:"email" json:"email"`
	Firstname     string             `bson:"first_name" json:"first_name"`
	Surname       string             `bson:"surname" json:"surname"`
	GoogleID      string             `bson:"google_id" json:"google_id"`
	Picture       string             `bson:"picture" json:"picture"`
	VerifiedEmail bool               `bson:"verified_email" json:"verified_email"`
	Role          string             `bson:"role" json:"role"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}
