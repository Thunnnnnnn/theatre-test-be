package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Seat struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	TheatreId    primitive.ObjectID  `bson:"theatre_id" json:"theatre_id"`
	Row          int                 `bson:"row" json:"row"`
	Number       int                 `bson:"number" json:"number"`
	Status       string              `bson:"status" json:"status"`
	IsHolded     bool                `bson:"is_holded" json:"is_holded"`
	HoldedById   *primitive.ObjectID `bson:"holded_by_id,omitempty" json:"holded_by_id,omitempty"`
	HoldedExpire *time.Time          `bson:"holded_expire,omitempty" json:"holded_expire"`
	IsBooked     bool                `bson:"is_booked" json:"is_booked"`
	BookedById   *primitive.ObjectID `bson:"booked_by_id,omitempty" json:"booked_by_id,omitempty"`
	CreatedAt    time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time           `bson:"updated_at" json:"updated_at"`
}

type SeatResponse struct {
	ID           primitive.ObjectID  `json:"id"`
	TheatreId    primitive.ObjectID  `json:"theatre_id"`
	Theatre      TheatreResponse     `json:"theatre,omitempty"`
	Row          int                 `json:"row"`
	Number       int                 `json:"number"`
	Status       string              `json:"status"`
	IsHolded     bool                `json:"is_holded"`
	HoldedById   *primitive.ObjectID `bson:"holded_by_id,omitempty" json:"holded_by_id,omitempty"`
	HoldedExpire *time.Time          `json:"holded_expire,omitempty"`
	IsBooked     bool                `json:"is_booked"`
	BookedById   *primitive.ObjectID `bson:"booked_by_id,omitempty" json:"booked_by_id,omitempty"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
}

type AdminSeatResponse struct {
	ID           primitive.ObjectID  `json:"id"`
	TheatreId    primitive.ObjectID  `json:"theatre_id"`
	Theatre      TheatreResponse     `json:"theatre,omitempty"`
	Row          int                 `json:"row"`
	Number       int                 `json:"number"`
	Status       string              `json:"status"`
	IsHolded     bool                `json:"is_holded"`
	HoldedById   *primitive.ObjectID `bson:"holded_by_id,omitempty" json:"holded_by_id,omitempty"`
	HoldedExpire *time.Time          `json:"holded_expire,omitempty"`
	IsBooked     bool                `json:"is_booked"`
	BookedById   *primitive.ObjectID `bson:"booked_by_id,omitempty" json:"booked_by_id,omitempty"`
	BookedBy     User                `json:"booked_by,omitempty"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
}
