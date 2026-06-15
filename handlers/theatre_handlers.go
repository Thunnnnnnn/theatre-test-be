package handlers

import (
	"log"
	"net/http"
	"theatre-test-api/helpers"
	"theatre-test-api/models"
	"theatre-test-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllTheatres(c *gin.Context) {
	log.Println("===== GetAllTheatres HIT =====")
	theatres, err := services.FindAllTheatres()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, theatres)
}

func GetTheatreById(c *gin.Context) {
	id := c.Param("id")

	theatre, err := services.FindTheatreById(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if theatre.ID.IsZero() {
		helpers.Error(c, http.StatusNotFound, "theatre not found")
		return
	}

	helpers.Success(c, http.StatusOK, theatre)
}

func CreateTheatre(c *gin.Context) {
	var theatre models.Theatre

	if err := c.ShouldBindJSON(&theatre); err != nil {
		helpers.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	createdTheatre, err := services.CreateTheatre(theatre)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusCreated, createdTheatre)
}

func UpdateTheatre(c *gin.Context) {
	id := c.Param("id")
	var theatre models.TheatreResponse

	if err := c.ShouldBindJSON(&theatre); err != nil {
		helpers.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	updatedTheatre, err := services.UpdateTheatre(id, theatre)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, updatedTheatre)
}

func DeleteTheatre(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteTheatre(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, gin.H{"message": "theatre deleted"})
}
