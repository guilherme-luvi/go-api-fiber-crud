package security

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/config"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestComparePassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)

	assert.NoError(t, err)

	err = ComparePassword(hashedPassword, password)
	assert.NoError(t, err)

	err = ComparePassword(hashedPassword, "wrongpassword")
	assert.Error(t, err)
}

func TestGenerateToken(t *testing.T) {
	config.SecretKey = []byte("mysecretkey")

	userID := uint(1)
	tokenString, err := GenerateToken(userID)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return []byte(config.SecretKey), nil
	})

	assert.NoError(t, err)
	assert.NotNil(t, token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		assert.Equal(t, true, claims["authorized"])
		assert.Equal(t, float64(userID), claims["userId"].(float64))
		exp := int64(claims["exp"].(float64))
		assert.True(t, exp > time.Now().Unix())
	} else {
		t.Fail()
	}
}
