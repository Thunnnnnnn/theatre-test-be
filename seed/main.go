package seed

import (
	"context"
	"fmt"
	"log"
	"theatre-test-api/database"
	"theatre-test-api/models"
	"theatre-test-api/services"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Seed() {
	if err := database.ConnectMongo(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	movies := []interface{}{
		bson.M{
			"name":       "Avengers",
			"created_at": time.Now(),
			"updated_at": time.Now(),
		},
		bson.M{
			"name":       "Spider-Man",
			"created_at": time.Now(),
			"updated_at": time.Now(),
		},
		bson.M{
			"name":       "Batman",
			"created_at": time.Now(),
			"updated_at": time.Now(),
		},
	}

	count, _ := database.MovieCollection.CountDocuments(
		ctx,
		bson.M{},
	)

	if count > 0 {
		log.Println("seed already exists")
		return
	}

	result, err := database.MovieCollection.InsertMany(ctx, movies)
	if err != nil {
		log.Fatal(err)
	}

	count, _ = database.TheatreCollection.CountDocuments(
		ctx,
		bson.M{},
	)

	if count > 0 {
		log.Println("seed already exists")
		return
	}

	for i, id := range result.InsertedIDs {
		movieID := id.(primitive.ObjectID)

		_, err := services.CreateTheatre(models.Theatre{
			TheatreName:    fmt.Sprintf("Theatre%d", i+1),
			Rows:           5,
			SeatsPerRow:    10,
			TotalSeats:     50,
			AvailableSeats: 50,
			MovieId:        movieID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})

		if err != nil {
			log.Println(err)
		}
	}

	user := bson.M{
		"full_name":      "John Doe",
		"email":          "johndoe@example.com",
		"first_name":     "John",
		"surname":        "Doe",
		"google_id":      "",
		"picture":        "",
		"verified_email": true,
		"role":           "ADMIN",
		"created_at":     time.Now(),
		"updated_at":     time.Now(),
	}

	count, _ = database.UserCollection.CountDocuments(
		ctx,
		bson.M{},
	)

	if count > 0 {
		log.Println("seed already exists")
		return
	}

	_, err = database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
}
