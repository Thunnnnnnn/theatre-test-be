package routes

import (
	"theatre-test-api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	users := r.Group("/auth")
	{
		users.POST("/google", handlers.GoogleLogin)
		users.POST("/admin", handlers.AdminLogin)
	}
}
