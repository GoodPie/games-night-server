package middleware

import (
	"context"
	"net/http"
	"strings"
	"firebase.google.com/go/auth"
)

func AuthMiddleware(app *auth.Client) func (httpHandler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			authHeader := request.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(responseWriter, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			idToken := strings.Replace(authHeader, "Bearer ", "", 1)
			token, err := app.VerifyIDToken(request.Context(), idToken)
			if err != nil {
				http.Error(responseWriter, "Invalid ID token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(request.Context(), "user", token)
			next.ServeHTTP(responseWriter, request.WithContext(ctx))
		})
	}
}