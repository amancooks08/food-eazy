package utils

import (
	"time"
	"os"
	jwt "github.com/dgrijalva/jwt-go"
)

var securiy_key = []byte(os.Getenv("SECURITY_KEY"))

func GenerateToken(email, role string) (token string, err error) {
	tokenExpirationTime := time.Now().Add(time.Minute * 30)
	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   tokenExpirationTime.Unix(),
	})
	token, err = tokenObject.SignedString(securiy_key)
	return token, err
}
