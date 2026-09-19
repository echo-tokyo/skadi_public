// Package usecase contains auth.UsecaseClient and auth.UsecaseMiddleware implementations.
package usecase

import (
	"fmt"

	"skadi/backend/config"
	"skadi/backend/internal/app/auth"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/user"
	"skadi/backend/internal/pkg/jwt"
	"skadi/backend/internal/pkg/password"
)

// Ensure UCClient implements interface.
var _ auth.UsecaseClient = (*UCClient)(nil)

// UCClient represents an auth usecase for client.
// It implements the [auth.UsecaseClient] interface.
type UCClient struct {
	cfg           *config.Config
	userRepoDB    user.RepositoryDB
	authRepoCache auth.RepositoryCache
	jwtBuilder    jwt.Builder
}

// NewUCClient returns a new instance of [UCClient].
func NewUCClient(cfg *config.Config, userRepoDB user.RepositoryDB,
	authRepoCache auth.RepositoryCache) *UCClient {

	return &UCClient{
		cfg:           cfg,
		userRepoDB:    userRepoDB,
		authRepoCache: authRepoCache,
		jwtBuilder: *jwt.NewBuilder(
			cfg.Auth.AccessToken.Secret,
			cfg.Auth.RefreshToken.Secret,
			jwt.WithAccessTTL(cfg.Auth.AccessToken.TTL),
			jwt.WithRefreshTTL(cfg.Auth.RefreshToken.TTL)),
	}
}

// LogIn returns authenticated user with token pair (access and refresh).
// Password is a raw (not hashed) password.
func (u *UCClient) LogIn(username string, passwd []byte) (*entity.UserWithToken, error) {
	// get userObj from DB with username
	userObj, err := u.userRepoDB.GetOneFull("username", username)
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}
	// schedule is unnecessary data in this usecase
	if userObj.Class != nil {
		userObj.Class.Schedule = nil
	}

	// check entered password is correct
	if !password.IsCorrect(passwd, userObj.Password) {
		return nil, auth.ErrInvalidPassword
	}

	// obtain token pair
	token, err := u.obtainTokenPair(userObj)
	if err != nil {
		return nil, err
	}
	return &entity.UserWithToken{
		User:  userObj,
		Token: token,
	}, nil
}

// LogOut sets given refresh token to blacklist.
func (u *UCClient) LogOut(refreshToken string) error {
	// put token to blacklist
	err := u.authRepoCache.SetTokenToBlacklist(refreshToken, u.cfg.Auth.RefreshToken.TTL)
	if err != nil {
		return fmt.Errorf("blacklist token: %w", err)
	}
	return nil
}

// ObtainAccess obtains a new access token using given user data.
func (u *UCClient) ObtainAccess(userClaims *entity.UserClaims) (*entity.Token, error) {
	// generate access token
	accessToken, err := u.jwtBuilder.ObtainAccess(userClaims)
	if err != nil {
		return nil, err
	}
	return &entity.Token{Access: accessToken}, nil
}

// obtainTokenPair returns obtained token pair (access and refresh) for user.
func (u *UCClient) obtainTokenPair(userObj *entity.User) (*entity.Token, error) {
	userClaims := &entity.UserClaims{
		ID:   userObj.ID,
		Role: userObj.Role,
	}
	// generate token pair
	accessToken, err := u.jwtBuilder.ObtainAccess(userClaims)
	if err != nil {
		return nil, fmt.Errorf("obtain access token: %w", err)
	}
	refreshToken, err := u.jwtBuilder.ObtainRefresh(userClaims)
	if err != nil {
		return nil, fmt.Errorf("obtain refresh token: %w", err)
	}
	return &entity.Token{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}
