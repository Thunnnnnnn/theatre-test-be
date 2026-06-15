package routes

import (
	"theatre-test-api/handlers"
	"theatre-test-api/middleware"

	"github.com/gin-gonic/gin"
)

func SeatRoutes(r *gin.RouterGroup) {
	seatGroup := r.Group("/seats")
	seatGroup.Use(middleware.AuthMiddleware())
	{
		seatGroup.GET("", handlers.GetAllSeats)
		seatGroup.GET("/:id", handlers.GetSeatById)
		seatGroup.POST("", handlers.CreateSeat)
		seatGroup.PUT("/:id", handlers.UpdateSeat)
		seatGroup.DELETE("/:id", handlers.DeleteSeat)
		seatGroup.GET("/theatre/:theatreId", handlers.GetSeatsByTheatreId)
		seatGroup.PUT("/booked/:id", handlers.UpdateBookedStatus)
		seatGroup.PUT("/hold/:id", handlers.UpdateIsHoldedStatus)
		seatGroup.GET("/user-booked", handlers.GetSeatUserByIsBookedAndIsHolded)
	}

	adminSeat := seatGroup.Group("")
	adminSeat.Use(middleware.AuthMiddleware())
	adminSeat.Use(middleware.RequireRole("ADMIN"))
	{
		adminSeat.GET("/booked", handlers.GetSeatsByIsBooked)
	}
}
