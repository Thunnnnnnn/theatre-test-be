package repositories

import (
	"context"
	"time"

	"theatre-test-api/database"
	"theatre-test-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllMovies() ([]models.Movie, error) {
	ctx := context.Background()

	if database.MovieCollection == nil {
		return nil, nil
	}

	cursor, err := database.MovieCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var movies []models.Movie
	if err := cursor.All(ctx, &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func FindMovieById(id primitive.ObjectID) (models.Movie, error) {
	ctx := context.Background()

	if database.MovieCollection == nil {
		return models.Movie{}, nil
	}

	var movie models.Movie
	err := database.MovieCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Movie{}, nil
		}
		return models.Movie{}, err
	}

	return movie, nil
}

func CreateMovie(movie models.Movie) (models.Movie, error) {
	ctx := context.Background()
	now := time.Now()

	if database.MovieCollection == nil {
		return models.Movie{}, nil
	}

	movie.CreatedAt = now
	movie.UpdatedAt = now

	result, err := database.MovieCollection.InsertOne(ctx, movie)
	if err != nil {
		return models.Movie{}, err
	}

	movie.ID = result.InsertedID.(primitive.ObjectID)
	return movie, nil
}

func UpdateMovie(id primitive.ObjectID, movie models.Movie) (models.Movie, error) {
	ctx := context.Background()

	if database.MovieCollection == nil {
		return models.Movie{}, nil
	}

	findMovie, err := FindMovieById(id)
	if err != nil {
		return models.Movie{}, err
	}

	movie.CreatedAt = findMovie.CreatedAt
	movie.UpdatedAt = time.Now()

	_, err = database.MovieCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": movie})
	if err != nil {
		return models.Movie{}, err
	}

	return FindMovieById(id)
}

func DeleteMovie(id primitive.ObjectID) error {
	ctx := context.Background()

	if database.MovieCollection == nil {
		return nil
	}
	_, err := database.MovieCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
