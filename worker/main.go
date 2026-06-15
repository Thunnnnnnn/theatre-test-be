package main

import (
	"context"
	"log"
	"theatre-test-api/database"
	"theatre-test-api/redisclient"
	"theatre-test-api/seat_release"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}

	redisclient.InitRedis()

	if err := database.ConnectMongo(); err != nil {
		log.Fatal("Mongo connection failed:", err)
	}

	releaser := seat_release.NewSeatReleaser(database.SeatCollection, redisclient.Rdb)

	log.Println("worker started...")

	for {
		err := releaser.ReleaseExpiredSeats(ctx)
		if err != nil {
			log.Println("error:", err)
		}

		time.Sleep(30 * time.Second)
	}
}
