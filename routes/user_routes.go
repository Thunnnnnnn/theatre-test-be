package routes

import (
	"theatre-test-api/handlers"
	"theatre-test-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("", handlers.GetAllUsers)
		users.GET("/profile", handlers.GetProfileUserByToken)
	}
}
