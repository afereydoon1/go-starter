package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestJWTService(t *testing.T) {
	jwtService := NewJWTService("test_secret")

	t.Run("Generate and parse JWT", func(t *testing.T) {
		userID := uint(123)
		token, err := jwtService.GenerateJWT(userID)
		assert.NoError(t, err, "GenerateJWT should not return error")
		assert.NotEmpty(t, token, "Token should not be empty")

		parsedID, err := jwtService.ParseJWT(token)
		assert.NoError(t, err, "ParseJWT should not return error")
		assert.Equal(t, userID, parsedID, "Parsed userID should match")
	})

	t.Run("Generate and parse Refresh Token", func(t *testing.T) {
		userID := uint(123)
		token, err := jwtService.GenerateRefreshToken(userID)
		assert.NoError(t, err, "GenerateRefreshToken should not return error")
		assert.NotEmpty(t, token, "Refresh token should not be empty")

		parsedID, err := jwtService.ParseJWT(token)
		assert.NoError(t, err, "ParseJWT should not return error for refresh token")
		assert.Equal(t, userID, parsedID, "Parsed userID should match")
	})

	t.Run("Invalid token", func(t *testing.T) {
		_, err := jwtService.ParseJWT("invalid.token.here")
		assert.Error(t, err, "ParseJWT should return error for invalid token")
	})

	t.Run("Blacklist token", func(t *testing.T) {
		token, _ := jwtService.GenerateJWT(123)
		BlacklistToken(token)
		assert.True(t, IsTokenBlacklisted(token), "Token should be blacklisted")

		_, err := jwtService.ParseJWT(token)
		assert.Error(t, err, "ParseJWT should return error for blacklisted token")
	})
}