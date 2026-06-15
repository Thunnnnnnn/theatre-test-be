package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Theatre struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TheatreName    string             `bson:"theatre_name" json:"theatre_name"`
	Rows           int                `bson:"rows" json:"rows"`
	SeatsPerRow    int                `bson:"seats_per_row" json:"seats_per_row"`
	TotalSeats     int                `bson:"total_seats" json:"total_seats"`
	AvailableSeats int                `bson:"available_seats" json:"available_seats"`
	MovieId        primitive.ObjectID `bson:"movie_id" json:"movie_id"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type TheatreResponse struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TheatreName    string             `bson:"theatre_name" json:"theatre_name"`
	Rows           int                `bson:"rows" json:"rows"`
	SeatsPerRow    int                `bson:"seats_per_row" json:"seats_per_row"`
	TotalSeats     int                `bson:"total_seats" json:"total_seats"`
	AvailableSeats int                `bson:"available_seats" json:"available_seats"`
	MovieId        primitive.ObjectID `bson:"movie_id" json:"movie_id"`
	Movie          Movie              `bson:"movie" json:"movie"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}
