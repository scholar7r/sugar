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

var j = mjwt.New[userClaims]("secret")

func TestJWT_Generate(t *testing.T) {
	claims := &mjwt.Claims[userClaims]{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
		Data: &userClaims{ID: 1, Name: "test"},
	}

	_, err := j.Generate(claims)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJWT_Parse(t *testing.T) {
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
