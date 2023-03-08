package auth

import (
	"testing"
	"time"

	"github.com/borud/tt/pkg/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	auth, err := NewJWT(JWTConfig{
		Issuer:          "https://service.example.com",
		Secret:          []byte("this is a secret"),
		DefaultValidity: time.Hour * 24 * 30,
		Audience:        []string{"https://service.example.com"},
	})
	assert.NoError(t, err)

	user := model.User{
		Username: "testuser1",
	}

	signedToken, err := auth.Create(user)
	assert.NoError(t, err)

	claims, err := auth.ParseAndVerify(signedToken)
	assert.NoError(t, err)
	assert.NotNil(t, claims)

	signedToken, err = auth.Create(user, (24 * 100 * time.Hour))
	assert.NoError(t, err)

	claims, err = auth.ParseAndVerify(signedToken)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
}

func TestInvalidToken(t *testing.T) {
	auth, err := NewJWT(JWTConfig{
		Issuer:          "foo",
		Secret:          []byte("this is a different secret"),
		DefaultValidity: time.Hour * 24 * 30,
		Audience:        []string{"https://service.example.com"},
	})
	assert.NoError(t, err)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJmb28iLCJzdWIiOiJ0ZXN0dXNlcjEiLCJhdWQiOlsiYXVkaWVuY2UiXSwiZXhwIjoxNjU4MTUzMzQwLCJuYmYiOjE2NTU1NjEzNDAsImlhdCI6MTY1NTU2MTM0MCwianRpIjoiNXc1bWJ1ZzM5MDNoYXIyM3kzYXhva2F5YyJ9.fOvzQpTbjUIchrs7C93yjFrert_z6GR__kan6XKa3PA"
	claims, err := auth.ParseAndVerify(token)
	assert.ErrorIs(t, err, jwt.ErrSignatureInvalid)
	assert.Nil(t, claims)
}
