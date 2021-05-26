package app

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if len(tokenString) == 0 {
			NewJSendJSONBuilder().
				Code(http.StatusUnauthorized).
				Message("Missing Authorization Header").
				Build().
				Send(w)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := VerifyToken(tokenString)
		if err != nil {
			NewJSendJSONBuilder().
				Code(http.StatusUnauthorized).
				Message("Error verifying JWT token:" + err.Error()).
				Build().
				Send(w)
			return
		}
		addr := claims.(jwt.MapClaims)["Address"].(string)
		r.Header.Set("address", addr)
		next.ServeHTTP(w, r)
	})
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(SECRET_JWT)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
