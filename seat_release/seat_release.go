package seat_release

import (
	"context"

	"theatre-test-api/repositories"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type SeatReleaser struct {
	db    *mongo.Collection
	redis *redis.Client
}

func NewSeatReleaser(db *mongo.Collection, redis *redis.Client) *SeatReleaser {
	return &SeatReleaser{
		db:    db,
		redis: redis,
	}
}

func (s *SeatReleaser) ReleaseExpiredSeats(ctx context.Context) error {
	repositories.UpdateReleaseHoldedSeat()

	return nil
}
