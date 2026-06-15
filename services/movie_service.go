package services

import (
	"theatre-test-api/models"
	"theatre-test-api/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllMovies() ([]models.Movie, error) {
	return repositories.FindAllMovies()
}

func FindMovieById(id string) (models.Movie, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Movie{}, err
	}

	return repositories.FindMovieById(objectID)
}

func CreateMovie(movie models.Movie) (models.Movie, error) {
	return repositories.CreateMovie(movie)
}

func UpdateMovie(id string, movie models.Movie) (models.Movie, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Movie{}, err
	}

	return repositories.UpdateMovie(objectID, movie)
}

func DeleteMovie(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return repositories.DeleteMovie(objectID)
}
