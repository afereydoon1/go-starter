package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "test123"
	hash, err := HashPassword(password)

	assert.NoError(t, err, "HashPassword should not return error")
	assert.NotEmpty(t, hash, "Hash should not be empty")
	assert.NotEqual(t, password, hash, "Hash should not be the same as plain password")
}

func TestCheckPassword(t *testing.T) {
	password := "test123"
	hash, _ := HashPassword(password)

	assert.True(t, CheckPassword(password, hash), "CheckPassword should return true for correct password")
	assert.False(t, CheckPassword("wrong", hash), "CheckPassword should return false for wrong password")
}