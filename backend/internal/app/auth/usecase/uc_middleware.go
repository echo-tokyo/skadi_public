package usecase

import (
	"fmt"

	"skadi/backend/config"
	"skadi/backend/internal/app/auth"
)

// Ensure UCMiddleware implements interface.
var _ auth.UsecaseMiddleware = (*UCMiddleware)(nil)

// UCMiddleware represents an auth usecase for middlewares.
// It implements the [auth.UsecaseMiddleware] interface.
type UCMiddleware struct {
	cfg           *config.Config
	authRepoCache auth.RepositoryCache
}

// NewUCMiddleware returns a new instance of [UCMiddleware].
func NewUCMiddleware(cfg *config.Config, authRepoCache auth.RepositoryCache) *UCMiddleware {
	return &UCMiddleware{
		cfg:           cfg,
		authRepoCache: authRepoCache,
	}
}

// BlockIfBlacklist returns error if the given token is in blacklist.
func (u *UCMiddleware) BlockIfBlacklist(token string) error {
	// search token in blacklist
	isBlacklisted, err := u.authRepoCache.TokenIsBlacklisted(token)
	if err != nil {
		return err
	}
	// return forbidden error if token is in blacklist
	if isBlacklisted {
		return fmt.Errorf("%w: token is in blacklist", auth.ErrForbidden)
	}
	return nil
}
