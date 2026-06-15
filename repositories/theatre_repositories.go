package repositories

import (
	"context"
	"log"
	"time"

	"theatre-test-api/database"
	"theatre-test-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllTheatres() ([]models.TheatreResponse, error) {
	ctx := context.Background()

	if database.TheatreCollection == nil {
		return nil, nil
	}

	cursor, err := database.TheatreCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var theatres []models.Theatre
	if err := cursor.All(ctx, &theatres); err != nil {
		return nil, err
	}

	var theatreResponses []models.TheatreResponse
	for _, theatre := range theatres {

		movie, err := FindMovieById(theatre.MovieId)
		if err != nil {
			log.Println("FindMovieById error:", err)
		}
		theatreResponses = append(theatreResponses, models.TheatreResponse{
			ID:             theatre.ID,
			TheatreName:    theatre.TheatreName,
			Rows:           theatre.Rows,
			SeatsPerRow:    theatre.SeatsPerRow,
			TotalSeats:     theatre.TotalSeats,
			AvailableSeats: theatre.AvailableSeats,
			MovieId:        theatre.MovieId,
			Movie:          movie,
			CreatedAt:      theatre.CreatedAt,
			UpdatedAt:      theatre.UpdatedAt,
		})
	}

	return theatreResponses, nil
}

func FindTheatreById(id primitive.ObjectID) (models.TheatreResponse, error) {
	ctx := context.Background()

	if database.TheatreCollection == nil {
		return models.TheatreResponse{}, nil
	}

	var theatre models.Theatre
	err := database.TheatreCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&theatre)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.TheatreResponse{}, nil
		}
		return models.TheatreResponse{}, err
	}

	var theatreResponse models.TheatreResponse

	movie, _ := FindMovieById(theatre.MovieId)
	theatreResponse = models.TheatreResponse{
		ID:             theatre.ID,
		TheatreName:    theatre.TheatreName,
		Rows:           theatre.Rows,
		SeatsPerRow:    theatre.SeatsPerRow,
		TotalSeats:     theatre.TotalSeats,
		AvailableSeats: theatre.AvailableSeats,
		MovieId:        theatre.MovieId,
		Movie:          movie,
		CreatedAt:      theatre.CreatedAt,
		UpdatedAt:      theatre.UpdatedAt,
	}

	return theatreResponse, nil
}

func CreateTheatre(theatre models.Theatre) (models.Theatre, error) {
	ctx := context.Background()

	if database.TheatreCollection == nil {
		return models.Theatre{}, nil
	}

	theatre.CreatedAt = time.Now()
	theatre.UpdatedAt = time.Now()

	theatre.TotalSeats = theatre.Rows * theatre.SeatsPerRow
	theatre.AvailableSeats = theatre.TotalSeats

	result, err := database.TheatreCollection.InsertOne(ctx, theatre)
	if err != nil {
		return models.Theatre{}, err
	}

	for i := 1; i <= theatre.Rows; i++ {
		for j := 1; j <= theatre.SeatsPerRow; j++ {
			seat := models.Seat{
				TheatreId: result.InsertedID.(primitive.ObjectID),
				Row:       i,
				Number:    j,
				IsHolded:  false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsBooked:  false,
				Status:    "available",
			}
			_, err := database.SeatCollection.InsertOne(ctx, seat)
			if err != nil {
				return models.Theatre{}, err
			}
		}
	}

	theatre.ID = result.InsertedID.(primitive.ObjectID)
	return theatre, nil
}

func UpdateTheatre(id primitive.ObjectID, theatre models.TheatreResponse) (models.TheatreResponse, error) {
	ctx := context.Background()

	if database.TheatreCollection == nil {
		return models.TheatreResponse{}, nil
	}

	findTheatre, err := FindTheatreById(id)
	if err != nil {
		return models.TheatreResponse{}, err
	}

	theatre.CreatedAt = findTheatre.CreatedAt
	theatre.UpdatedAt = time.Now()

	_, err = database.TheatreCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"theatre_name":    theatre.TheatreName,
		"rows":            theatre.Rows,
		"seats_per_row":   theatre.SeatsPerRow,
		"total_seats":     theatre.TotalSeats,
		"available_seats": theatre.AvailableSeats,
		"movie_id":        theatre.MovieId,
		"created_at":      theatre.CreatedAt,
		"updated_at":      theatre.UpdatedAt,
	}})
	if err != nil {
		return models.TheatreResponse{}, err
	}

	return FindTheatreById(id)
}

func DeleteTheatre(id primitive.ObjectID) error {
	ctx := context.Background()

	if database.TheatreCollection == nil {
		return nil
	}
	_, err := database.TheatreCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
