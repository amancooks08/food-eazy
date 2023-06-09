package utils

import (
	"auth-service/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Run("Generate token with valid user", func(t *testing.T) {
		email, role := "test1@mail.com", "USER"
		var id uint = 1
		token, err := GenerateToken(id, email, role)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("Generate token with empty user", func(t *testing.T) {
		email, role := "", "USER"
		var id uint = 1
		token, err := GenerateToken(id, email, role)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrEmptyField.Error(), err.Error())
		assert.Empty(t, token)
	})
	t.Run("Generate token with empty role", func(t *testing.T) {
		email, role := "test1@gmail.com", ""
		var id uint = 1
		token, err := GenerateToken(id, email, role)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrEmptyField.Error(), err.Error())
		assert.Empty(t, token)
	})

	t.Run("Generate token with invalid role", func(t *testing.T) {
		email, role := "testingg@mail.com", "ADMINN"
		var id uint = 1
		token, err := GenerateToken(id, email, role)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidRole.Error(), err.Error())
		assert.Empty(t, token)
	})
}

func TestValidateToken(t *testing.T) {
	t.Run("Validate token with valid token", func(t *testing.T) {
		email, role := "test1@mail.com", "USER"
		var id uint = 1
		token, err := GenerateToken(id, email, role)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, err := ValidateToken(token)
		assert.NoError(t, err)
		assert.Equal(t, claims["email"], email)
		assert.Equal(t, claims["role"], role)
		userID, ok := claims["user_id"].(float64)
		assert.True(t, ok)

		user := uint(userID)
		assert.EqualValues(t, id, user)
	})
}


