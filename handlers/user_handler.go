package handlers

import (
	"net/http"
	"theatre-test-api/helpers"
	"theatre-test-api/services"

	"github.com/gin-gonic/gin"
)

func GetProfileUserByToken(c *gin.Context) {
	userID, _ := c.Get("userId")

	user, err := services.FindProfileByToken(userID.(string))
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	users, err := services.FindAllUsers()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := services.FindUserById(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, user)
}
