package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	_queryTimeout = 2 * time.Second // ctx timeout for cache queries
)

// Ensure Redis implements interface.
var _ Storage = (*Redis)(nil)

// Redis is a cache key-value storage based on Redis.
// Redis implements the [Storage] interface.
type Redis struct {
	client *redis.Client
}

// NewRedis returns a new instance of [Redis].
func NewRedis(connString string) (*Redis, error) {
	// create new client
	client := redis.NewClient(&redis.Options{
		Addr: connString,
		DB:   0,
	})
	// redis requests context
	ctx, cancel := context.WithTimeout(context.Background(), _queryTimeout)
	defer cancel()

	// check connection
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("connect to redis: %w", err)
	}
	return &Redis{
		client: client,
	}, nil
}

// Get gets the value for the given key.
func (s *Redis) Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), _queryTimeout)
	defer cancel()
	value, err := s.client.Get(ctx, key).Bytes()

	// if NOT "Not found" error
	if err != nil && !redis.HasErrorPrefix(err, "redis: nil") {
		return nil, fmt.Errorf("get value: %w", err)
	}
	return value, nil
}

// Set stores the given value for the given key along
// with an expiration value, 0 means no expiration.
func (s *Redis) Set(key string, val []byte, exp time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), _queryTimeout)
	defer cancel()
	if err := s.client.Set(ctx, key, val, exp).Err(); err != nil {
		return fmt.Errorf("set value: %w", err)
	}
	return nil
}

// Delete deletes the value for the given key.
func (s *Redis) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), _queryTimeout)
	defer cancel()
	if err := s.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("delete value: %w", err)
	}
	return nil
}

// Reset deletes all keys.
func (s *Redis) Reset() error {
	ctx, cancel := context.WithTimeout(context.Background(), _queryTimeout)
	defer cancel()
	if err := s.client.FlushAll(ctx).Err(); err != nil {
		return fmt.Errorf("reset storage: %w", err)
	}
	return nil
}

// Close closes the redis client.
func (s *Redis) Close() error {
	if err := s.client.Close(); err != nil {
		return fmt.Errorf("close client: %w", err)
	}
	return nil
}
