package handlers

import (
	"context"
	"fmt"
	"net/http"
	"theatre-test-api/helpers"
	"theatre-test-api/lock"
	"theatre-test-api/models"
	"theatre-test-api/redisclient"
	"theatre-test-api/services"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllSeats(c *gin.Context) {
	seats, err := services.FindAllSeats()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, seats)
}

func GetSeatById(c *gin.Context) {
	id := c.Param("id")

	seat, err := services.FindSeatById(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if seat.ID.IsZero() {
		helpers.Error(c, http.StatusNotFound, "seat not found")
		return
	}

	helpers.Success(c, http.StatusOK, seat)
}

func CreateSeat(c *gin.Context) {
	var seat models.Seat

	if err := c.ShouldBindJSON(&seat); err != nil {
		helpers.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	createdSeat, err := services.CreateSeat(seat)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusCreated, createdSeat)
}

func UpdateSeat(c *gin.Context) {
	id := c.Param("id")
	var seat models.Seat

	if err := c.ShouldBindJSON(&seat); err != nil {
		helpers.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	updatedSeat, err := services.UpdateSeat(id, seat)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, updatedSeat)
}

func DeleteSeat(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteSeat(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, gin.H{"message": "seat deleted"})
}

func GetSeatsByTheatreId(c *gin.Context) {
	theatreId := c.Param("theatreId")

	seats, err := services.FindSeatsByTheatreId(theatreId)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, seats)
}

func GetSeatsByIsBooked(c *gin.Context) {
	keyword := c.Query("keyword")

	seats, err := services.FindSeatsByIsBooked(keyword)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, seats)
}

func UpdateBookedStatus(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userId")

	updatedSeat, err := services.UpdateBookedStatus(id, userID.(string))
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, updatedSeat)
}

func UpdateIsHoldedStatus(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	userID, _ := c.Get("userId")
	rl := lock.New(redisclient.Rdb)

	lockKey := fmt.Sprintf(`seat:%s`, id)
	ttl := 10 * time.Second

	value, ok, err := rl.Acquire(ctx, lockKey, ttl)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(409, gin.H{"message": "seat locked"})
		return
	}

	fmt.Println("lock acquired with value:", value)

	updatedSeat, err := services.UpdateIsHoldedStatus(id, userID.(string))
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, updatedSeat)

}

func GetSeatUserByIsBookedAndIsHolded(c *gin.Context) {
	// id := c.Param("id")
	userIDRaw, _ := c.Get("userId")

	fmt.Printf("userID raw: %#v\n", userIDRaw)

	userID, err := primitive.ObjectIDFromHex(userIDRaw.(string))
	if err != nil {
		helpers.Error(c, 400, "invalid user id")
		return
	}

	seats, err := services.GetSeatUserByIsBookedAndIsHolded(userID)

	if err != nil {
		fmt.Println("Error:", err)
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, gin.H{"seats": seats})

}

func ReleaseHoldedSeat(c *gin.Context) {

	err := services.UpdateReleaseHoldedSeat()

	if err != nil {
		fmt.Println("Error:", err)
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, gin.H{"message": "holded seats released"})
}
