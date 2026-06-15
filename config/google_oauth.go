package config

import (
	"errors"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig *oauth2.Config

func GoogleOAuthInit() error {
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	if GoogleOAuthConfig.ClientID == "" {
		return errors.New("GOOGLE_CLIENT_ID not found")
	}

	if GoogleOAuthConfig.ClientSecret == "" {
		return errors.New("GOOGLE_CLIENT_SECRET not found")
	}

	return nil
}
