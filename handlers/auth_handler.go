package handlers

import (
	"fmt"
	"net/http"

	"theatre-test-api/database"
	"theatre-test-api/helpers"
	"theatre-test-api/models"
	"theatre-test-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type GoogleLoginRequest struct {
	Credential string `json:"credential"`
}

func GoogleLogin(c *gin.Context) {
	var req GoogleLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	user, token, err := services.GoogleLogin(req.Credential)
	if err != nil {
		helpers.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, gin.H{
		"accessToken": token,
		"user":        user,
	})
}

func AdminLogin(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, "invalid request")
		return
	}
	username := req.Username
	password := req.Password

	fmt.Println("username =", username, "password =", password)

	if username != "admin" || password != "admin" {
		helpers.Error(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	user := models.User{}
	_ = database.UserCollection.FindOne(ctx, bson.M{"role": "ADMIN"}).Decode(&user)

	token, err := helpers.GenerateJWT(user.ID.Hex(), user)

	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, gin.H{
		"accessToken": token,
		"user": models.User{
			FullName:      "Admin",
			Email:         "admin@mail.com",
			Role:          "ADMIN",
			Firstname:     "Admin",
			Surname:       "Admin",
			GoogleID:      "",
			VerifiedEmail: true,
			Picture:       "",
		},
	})
}
