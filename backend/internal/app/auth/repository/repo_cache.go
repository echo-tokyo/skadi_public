package repository

import (
	"fmt"
	"time"

	"skadi/backend/config"
	"skadi/backend/internal/app/auth"
	"skadi/backend/internal/pkg/cache"
)

const _blacklistPrefix = "token:blacklisted:" // key prefix for blacklisted token values

var _blacklisted = []byte("1") // value for blacklisted tokens

// Ensure RepoCache implements interface.
var _ auth.RepositoryCache = (*RepoCache)(nil)

// RepoCache is an auth cache repo.
// It implements the [auth.RepoCache] interface.
type RepoCache struct {
	cfg          *config.Config
	cacheStorage cache.Storage
}

// NewRepoCache returns a new instance of [RepoCache].
func NewRepoCache(cfg *config.Config, cacheStorage cache.Storage) *RepoCache {
	return &RepoCache{
		cfg:          cfg,
		cacheStorage: cacheStorage,
	}
}

// Set token to blacklist expiring after exp duration.
func (r *RepoCache) SetTokenToBlacklist(token string, exp time.Duration) error {
	err := r.cacheStorage.Set(_blacklistPrefix+token, _blacklisted, exp)
	if err != nil {
		return fmt.Errorf("set to cache: %w", err)
	}
	return nil
}

// TokenIsBlacklisted returns true, if given token is blacklisted.
func (r *RepoCache) TokenIsBlacklisted(token string) (bool, error) {
	val, err := r.cacheStorage.Get(_blacklistPrefix + token)
	if err != nil {
		return false, fmt.Errorf("get from cache: %w", err)
	}
	return len(val) > 0, nil
}
