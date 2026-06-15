package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		AuthRoutes(api)
		TheatreRoutes(api)
		MovieRoutes(api)
		SeatRoutes(api)
		UserRoutes(api)
	}
}
