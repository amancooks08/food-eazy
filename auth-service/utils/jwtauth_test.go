package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Run("Generate token with valid user", func(t *testing.T) {
		email, role := "test1@mail.com", "USER"
		token, err := GenerateToken(email, role)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}

func TestValidateToken(t *testing.T) {
	t.Run("Validate token with valid token", func(t *testing.T) {
		email, role := "test1@mail.com", "USER"
		token, err := GenerateToken(email, role)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, err := ValidateToken(token)
		assert.NoError(t, err)
		assert.Equal(t, claims["email"], email)
		assert.Equal(t, claims["role"], role)
	})
}


