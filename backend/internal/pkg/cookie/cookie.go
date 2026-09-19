// Package cookie provides builder to create and clear cookies.
package cookie

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

// expiration duration to clear cookie
const _clearCookieTTL = -(24 * time.Hour) // nolint:mnd // day duration

// Builder represents a cookie builder for cookie management.
type Builder struct {
	ttl      time.Duration
	path     string
	secure   bool
	httpOnly bool
	sameSite string
}

// Option represents an option for [Builder] initializing.
type Option func(*Builder)

// NewBuilder returns a new instance of [Builder].
func NewBuilder(ttl time.Duration, options ...Option) *Builder {
	builder := &Builder{
		ttl:      ttl,
		path:     "/",
		secure:   false,
		httpOnly: false,
		sameSite: "",
	}

	// apply all options to customize Builder
	for _, opt := range options {
		opt(builder)
	}
	return builder
}

// WithPath sets path for cookie.
func WithPath(path string) Option {
	return func(b *Builder) {
		b.path = path
	}
}

// WithSecure sets secure for cookie.
func WithSecure(secure bool) Option {
	return func(b *Builder) {
		b.secure = secure
	}
}

// WithHTTPOnly sets HTTP-inly for cookie.
func WithHTTPOnly(httpOnly bool) Option {
	return func(b *Builder) {
		b.httpOnly = httpOnly
	}
}

// WithSameSite sets same site for cookie.
// Supported values: Strict, Lax.
func WithSameSite(rawSameSite string) Option {
	sameSite := rawSameSite
	if rawSameSite != "Strict" && rawSameSite != "Lax" {
		sameSite = ""
	}
	return func(b *Builder) {
		b.sameSite = sameSite
	}
}

// Create creates a new cookie with parameters from builder.
func (b *Builder) Create(key, value string) *fiber.Cookie {
	cookie := b.emptyCookie()
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = cookie.Expires.Add(b.ttl)
	return cookie
}

// Clear creates a new clear cookie with parameters from builder.
func (b *Builder) Clear(key string) *fiber.Cookie {
	cookie := b.emptyCookie()
	cookie.Name = key
	cookie.Expires = cookie.Expires.Add(_clearCookieTTL)
	return cookie
}

// emptyCookie create an empty cookie with builder parameters.
func (b *Builder) emptyCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "",
		Value:    "",
		Path:     b.path,
		HTTPOnly: b.httpOnly,
		Secure:   b.secure,
		SameSite: b.sameSite,
		Expires:  time.Now().UTC(),
	}
}
