package services

import (
	"theatre-test-api/models"
	"theatre-test-api/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllTheatres() ([]models.TheatreResponse, error) {
	return repositories.FindAllTheatres()
}

func FindTheatreById(id string) (models.TheatreResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.TheatreResponse{}, err
	}

	return repositories.FindTheatreById(objectID)
}

func CreateTheatre(theatre models.Theatre) (models.Theatre, error) {
	return repositories.CreateTheatre(theatre)
}

func UpdateTheatre(id string, theatre models.TheatreResponse) (models.TheatreResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.TheatreResponse{}, err
	}

	return repositories.UpdateTheatre(objectID, theatre)
}

func DeleteTheatre(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return repositories.DeleteTheatre(objectID)
}
