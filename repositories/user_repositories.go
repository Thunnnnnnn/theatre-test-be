package repositories

import (
	"context"
	"log"

	"theatre-test-api/database"
	"theatre-test-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllUsers() ([]models.User, error) {
	ctx := context.Background()

	if database.UserCollection == nil {
		return nil, nil
	}

	cursor, err := database.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func FindUserByGoogleId(googleId string) (models.User, error) {
	ctx := context.Background()

	if database.UserCollection == nil {
		log.Fatal("UserCollection is not initialized")
		return models.User{}, nil
	}
	var user models.User
	err := database.UserCollection.FindOne(ctx, bson.M{"google_id": googleId}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, nil
	}

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	ctx := context.Background()

	if database.UserCollection == nil {
		log.Fatal("UserCollection is not initialized")
		return models.User{}, nil
	}

	result, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		panic(err)
		// return models.User{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func FindProfileByToken(userId string) (models.User, error) {
	ctx := context.Background()

	if database.UserCollection == nil {
		log.Fatal("UserCollection is not initialized")
		return models.User{}, nil
	}

	var user models.User
	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return models.User{}, err
	}
	err = database.UserCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, nil
	}

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func FindUserById(id primitive.ObjectID) (models.User, error) {
	ctx := context.Background()

	if database.UserCollection == nil {
		log.Fatal("UserCollection is not initialized")
		return models.User{}, nil
	}

	var user models.User
	err := database.UserCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, nil
	}

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
