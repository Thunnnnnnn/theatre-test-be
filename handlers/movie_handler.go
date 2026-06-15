package handlers

import (
	"net/http"
	"theatre-test-api/helpers"
	"theatre-test-api/models"
	"theatre-test-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(c *gin.Context) {
	movies, err := services.FindAllMovies()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, movies)
}

func GetMovieById(c *gin.Context) {
	id := c.Param("id")

	movie, err := services.FindMovieById(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if movie.ID.IsZero() {
		helpers.Error(c, http.StatusNotFound, "movie not found")
		return
	}

	helpers.Success(c, http.StatusOK, movie)
}

func CreateMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		helpers.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	createdMovie, err := services.CreateMovie(movie)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusCreated, createdMovie)
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		helpers.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	updatedMovie, err := services.UpdateMovie(id, movie)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, updatedMovie)
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteMovie(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, gin.H{"message": "movie deleted"})
}
