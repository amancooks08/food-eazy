package utils

import (
	"time"
	"os"
	jwt "github.com/dgrijalva/jwt-go"
	"auth-service/errors"
)

var security_key = []byte(os.Getenv("SECURITY_KEY"))

func GenerateToken(id uint, email, role string) (token string, err error) {
	if len(email) == 0 || len(role) == 0 {
		return "", errors.ErrEmptyField
	}
	if role != "ADMIN" && role != "USER" {
		return "", errors.ErrInvalidRole
	}
	tokenExpirationTime := time.Now().Add(time.Minute * 30)
	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email": email,
		"role":  role,
		"exp":   tokenExpirationTime.Unix(),
	})
	token, err = tokenObject.SignedString(security_key)
	return token, err
}


func ValidateToken(token string) (claims jwt.MapClaims, err error) {
	tokenObject, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(security_key), nil
	})

	if err != nil {
		return nil, errors.ErrUnauthorized
	}

	if !tokenObject.Valid {
		return nil, errors.ErrUnauthorized
	}

	claims, ok := tokenObject.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.ErrUnauthorized
	}

	return claims, nil
}