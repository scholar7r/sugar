package jwt_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	mjwt "github.com/scholar7r/sugar/jwt"
)

type userClaims struct {
	ID   int64
	Name string
}

var j = mjwt.New[userClaims]("secret", jwt.SigningMethodHS256)

func TestJWT_GenerateAndParse(t *testing.T) {
	token, err := j.Generate(&mjwt.Claims[userClaims]{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
		Data: &userClaims{ID: 1, Name: "test"},
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = j.Parse(token)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJWT_ParseNil(t *testing.T) {
	token, err := j.Generate(nil)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := j.Parse(token)
	if err != nil {
		t.Fatal(err)
	}

	if parsed.Data != nil {
		t.Fatalf("expected Data to be nil, got %+v", parsed.Data)
	}
}

func TestJWT_ParseInvalid(t *testing.T) {
	parsed, err := j.Parse("invalid token")
	if err == nil {
		t.Fatal("expected error for invalid token, got nil")
	}

	if parsed != nil {
		t.Fatalf("expected parsed to be nil, got %+v", parsed)
	}
}

func TestJWT_WrongSigningMethod(t *testing.T) {
	j1 := mjwt.New[userClaims]("secret", jwt.SigningMethodHS256)
	j2 := mjwt.New[userClaims]("secret", jwt.SigningMethodHS384)

	token, err := j1.Generate(&mjwt.Claims[userClaims]{
		Data: &userClaims{ID: 1},
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = j2.Parse(token)
	if err == nil {
		t.Fatal("expected signing method mismatch error")
	}
}
