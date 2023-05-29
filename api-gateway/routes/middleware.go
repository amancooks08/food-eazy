package routes

import (
	"api-gateway/clients"
	"api-gateway/domain"
	"context"
	"net/http"
	"os"
	"strings"
	proto "api-gateway/proto/auth"
	"github.com/dgrijalva/jwt-go"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		header := req.Header.Get("Authorization")

		// Check if the header is missing or invalid
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			http.Error(rw, "Unauthorized", http.StatusForbidden)
			return
		}

		// Parse the JWT token from the header
		token, err := jwt.Parse(strings.TrimPrefix(header, "Bearer "), func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			// Set the secret key for the token
			return []byte(os.Getenv("SECURITY_KEY")), nil
		})

		// Check if there was an error parsing the token
		if err != nil {
			http.Error(rw, "Unauthorized", http.StatusForbidden)
			return
		}

		// Check if the token is valid and has not expired
		if !token.Valid {
			http.Error(rw, "Unauthorized", http.StatusForbidden)
			return
		}

		// Check if the user has the "venue_owner" role to access venue_owner routes
		if strings.HasPrefix(req.URL.Path, "/admin") {
			status, err := clients.AuthServiceClient.ValidateToken(req.Context(), &proto.ValidateTokenRequest{
				Token: token.Raw,
				Role: domain.RoleMap["ADMIN"],
			})
			if err != nil {
				http.Error(rw, "Unauthorized", http.StatusForbidden)
				return
			}
			if status.StatusCode != http.StatusOK {
				http.Error(rw, "Unauthorized", http.StatusForbidden)
				return
			}
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx := context.WithValue(req.Context(), "id", claims["user_id"])
		req = req.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(rw, req)
	})
}
