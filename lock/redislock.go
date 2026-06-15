package lock

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisLock struct {
	client *redis.Client
}

func New(client *redis.Client) *RedisLock {
	return &RedisLock{client: client}
}

func (l *RedisLock) Acquire(ctx context.Context, key string, ttl time.Duration) (string, bool, error) {
	value := uuid.New().String()

	ok, err := l.client.SetNX(ctx, key, value, ttl).Result()
	if err != nil {
		return "", false, err
	}

	return value, ok, nil
}

var unlockScript = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return 0
end
`

func (l *RedisLock) Release(ctx context.Context, key, value string) error {
	_, err := l.client.Eval(ctx, unlockScript, []string{key}, value).Result()
	return err
}
