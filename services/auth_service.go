package services

import (
	"context"
	"os"
	"theatre-test-api/helpers"
	"theatre-test-api/models"

	"google.golang.org/api/idtoken"
)

func GoogleLogin(credential string) (*models.User, string, error) {

	payload, err := idtoken.Validate(
		context.Background(),
		credential,
		os.Getenv("GOOGLE_CLIENT_ID"),
	)

	if err != nil {
		return nil, "", err
	}

	googleId := payload.Claims["sub"].(string)
	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)
	picture := payload.Claims["picture"].(string)
	firstName := payload.Claims["given_name"].(string)
	surname := payload.Claims["family_name"].(string)
	verifiedEmail := payload.Claims["email_verified"].(bool)

	if !verifiedEmail {
		return nil, "", err
	}

	findedUser, err := FindOrCreateGoogleUser(models.User{
		Email:         email,
		Firstname:     firstName,
		Surname:       surname,
		FullName:      name,
		Picture:       picture,
		VerifiedEmail: verifiedEmail,
		GoogleID:      googleId,
	})

	if err != nil {
		return nil, "", err
	}

	jwtToken, err := helpers.GenerateJWT(findedUser.ID.Hex(), models.User{
		FullName:      name,
		Email:         email,
		Firstname:     firstName,
		Surname:       surname,
		Picture:       picture,
		GoogleID:      googleId,
		VerifiedEmail: verifiedEmail,
		Role:          findedUser.Role,
		ID:            findedUser.ID,
	})

	if err != nil {
		return nil, "", err
	}

	return &models.User{
		FullName:      name,
		Email:         email,
		Firstname:     firstName,
		Surname:       surname,
		Picture:       picture,
		GoogleID:      googleId,
		VerifiedEmail: verifiedEmail,
		Role:          findedUser.Role,
		ID:            findedUser.ID,
	}, jwtToken, nil
}
