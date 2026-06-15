package main

import (
	"log"
	"theatre-test-api/config"
	"theatre-test-api/database"
	"theatre-test-api/kafka"
	"theatre-test-api/redisclient"
	"theatre-test-api/routes"
	"theatre-test-api/seed"
	"theatre-test-api/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	redisclient.InitRedis()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"*",
		},
	}))

	if err := database.ConnectMongo(); err != nil {
		log.Fatal("Mongo connection failed:", err)
	}

	if err := config.GoogleOAuthInit(); err != nil {
		log.Fatal("Google OAuth initialization failed:", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin",
		})
	})

	r.Use(func(c *gin.Context) {
		log.Println("REQUEST:", c.Request.Method, c.Request.URL.Path)

		c.Next()

		log.Println("STATUS:", c.Writer.Status())
	})

	seed.Seed()

	routes.SetupRoutes(r)

	go kafka.StartConsumer()

	r.GET("/ws", websocket.Handler)

	r.Run(":8080")
}
