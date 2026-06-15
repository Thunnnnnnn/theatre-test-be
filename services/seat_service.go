package services

import (
	"theatre-test-api/models"
	"theatre-test-api/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllSeats() ([]models.Seat, error) {
	return repositories.FindAllSeats()
}

func FindSeatById(id string) (models.Seat, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Seat{}, err
	}

	return repositories.FindSeatById(objectID)
}

func CreateSeat(seat models.Seat) (models.Seat, error) {
	return repositories.CreateSeat(seat)
}

func UpdateSeat(id string, seat models.Seat) (models.Seat, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Seat{}, err
	}

	return repositories.UpdateSeat(objectID, seat)
}

func DeleteSeat(id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return repositories.DeleteSeat(objectID)
}

func FindSeatsByTheatreId(theatreId string) ([]models.Seat, error) {
	objectID, err := primitive.ObjectIDFromHex(theatreId)

	if err != nil {
		return nil, err
	}

	return repositories.FindSeatsByTheatreId(objectID)
}

func FindSeatsByIsBooked(keyword string) ([]models.AdminSeatResponse, error) {
	return repositories.FindSeatsByIsBooked(keyword)
}

func UpdateBookedStatus(id string, userID string) (models.Seat, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Seat{}, err
	}

	return repositories.UpdateBookedStatus(objectID, userID)
}

func UpdateIsHoldedStatus(id string, userID string) (models.Seat, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Seat{}, err
	}

	return repositories.UpdateIsHoldedStatus(objectID, userID)
}

func GetSeatUserByIsBookedAndIsHolded(userID primitive.ObjectID) ([]models.SeatResponse, error) {
	return repositories.GetSeatUserByIsBookedAndIsHolded(userID)
}

func UpdateReleaseHoldedSeat() error {
	return repositories.UpdateReleaseHoldedSeat()
}
