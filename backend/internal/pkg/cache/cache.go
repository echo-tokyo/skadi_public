// Package cache provides storage interface for
// caching data and its implementations.
package cache

import (
	"time"

	_ "github.com/gofiber/fiber/v2" // imported for comment hint
)

// Storage represents a cache key-value storage.
// This interface like a [fiber.Storage] interface.
type Storage interface {
	// Get gets the value for the given key.
	Get(key string) ([]byte, error)
	// Set stores the given value for the given key along
	// with an expiration value, 0 means no expiration.
	Set(key string, val []byte, exp time.Duration) error
	// Delete deletes the value for the given key.
	Delete(key string) error
	// Reset deletes all keys.
	Reset() error
	// Close closes the redis client.
	Close() error
}
