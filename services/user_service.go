package services

import (
	"theatre-test-api/models"
	"theatre-test-api/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindOrCreateGoogleUser(google models.User) (models.User, error) {
	googleId := google.GoogleID

	user, err := repositories.FindUserByGoogleId(googleId)
	if err != nil {
		panic("err:" + err.Error())
	}

	if !user.ID.IsZero() {
		return user, nil
	}

	newUser := models.User{
		Email:         google.Email,
		Firstname:     google.Firstname,
		Surname:       google.Surname,
		FullName:      google.FullName,
		Picture:       google.Picture,
		VerifiedEmail: google.VerifiedEmail,
		GoogleID:      googleId,
		Role:          "USER",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return repositories.CreateUser(newUser)

}

func FindAllUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}

func FindProfileByToken(userID string) (models.User, error) {
	return repositories.FindProfileByToken(userID)
}

func FindUserById(id string) (models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	return repositories.FindUserById(objID)
}
