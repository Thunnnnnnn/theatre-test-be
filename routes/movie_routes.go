package routes

import (
	"theatre-test-api/handlers"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(r *gin.RouterGroup) {
	movieGroup := r.Group("/movies")
	{
		movieGroup.GET("", handlers.GetAllMovies)
		movieGroup.GET("/:id", handlers.GetMovieById)
		movieGroup.POST("", handlers.CreateMovie)
		movieGroup.PUT("/:id", handlers.UpdateMovie)
		movieGroup.DELETE("/:id", handlers.DeleteMovie)
	}
}
