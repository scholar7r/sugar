// Package jwt provides a minimal, type-safe wrapper around github.com/golang-jwt/jwt
// with support for generic custom claims data.
package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrIsNil              = errors.New("token is nil")
	ErrNotValid           = errors.New("token not valid")
	ErrClaimsTypeMismatch = errors.New("claims type mismatch")
)

// JWT represents a JWT handler using a shared secret and generic claims data.
type JWT[T any] struct {
	Secret []byte
}

// New creates a new JWT instance using the given secret string.
func New[T any](secret string) *JWT[T] {
	return &JWT[T]{
		Secret: []byte(secret),
	}
}

// Claims represents JWT claims with embedded standard registered claims
// and an optional generic data payload.
type Claims[T any] struct {
	jwt.RegisteredClaims

	Data *T `json:"data,omitempty"`
}

// Generate creates and signs a JWT token using the provided claims.
func (x *JWT[T]) Generate(claims *Claims[T]) (string, error) {
	v := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return v.SignedString(x.Secret)
}

// Parse parses and validates a JWT token string and returns the claims
// if the token is valid.
func (x *JWT[T]) Parse(tokenString string) (*Claims[T], error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims[T]{}, func(_ *jwt.Token) (any, error) {
		return x.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, ErrIsNil
	}

	if !token.Valid {
		return nil, ErrNotValid
	}

	claims, ok := token.Claims.(*Claims[T])
	if !ok {
		return nil, ErrClaimsTypeMismatch
	}

	return claims, nil
}
