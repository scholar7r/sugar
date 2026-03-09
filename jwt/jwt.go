// SPDX-License-Identifier: GPLv2
// Copyright (c) 2026 scholar7r.

// Package jwt provides a minimal, type-safe wrapper around github.com/golang-jwt/jwt
// with support for generic custom claims data.
package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// JWT represents a JWT handler using a shared secret and generic claims data.
type JWT[T any] struct {
	secret []byte
	method jwt.SigningMethod
}

// New creates a new JWT instance using the given secret string.
func New[T any](secret string, method jwt.SigningMethod) *JWT[T] {
	return &JWT[T]{
		secret: []byte(secret),
		method: method,
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
	v := jwt.NewWithClaims(x.method, claims)

	return v.SignedString(x.secret)
}

// Parse parses and validates a JWT token string and returns the claims
// if the token is valid.
func (x *JWT[T]) Parse(tokenString string) (*Claims[T], error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims[T]{}, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != x.method.Alg() {
			return nil, jwt.ErrTokenSignatureInvalid
		}

		return x.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(*Claims[T])

	return claims, nil
}
