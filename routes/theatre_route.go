package routes

import (
	"theatre-test-api/handlers"

	"github.com/gin-gonic/gin"
)

func TheatreRoutes(r *gin.RouterGroup) {
	theatreGroup := r.Group("/theatres")
	{
		theatreGroup.GET("", handlers.GetAllTheatres)
		theatreGroup.GET("/:id", handlers.GetTheatreById)
		theatreGroup.POST("", handlers.CreateTheatre)
		theatreGroup.PUT("/:id", handlers.UpdateTheatre)
		theatreGroup.DELETE("/:id", handlers.DeleteTheatre)
	}
}
