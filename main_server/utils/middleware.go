package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}


type MiddlewareBody struct {
	Protected   bool
	Handlerfunc http.Handler
}

func Protected(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ErrorResponse(w, r, http.StatusUnauthorized, "Authorization header missing")
			return
		}

		// Ensure it's a Bearer token
		const bearerPrefix = "Bearer "
		if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			ErrorResponse(w, r, http.StatusUnauthorized, "Invalid token format")
			return
		}

		// Extract the actual token
		token := authHeader[len(bearerPrefix):]

		claims, err := DecodeToken(token)
		if err != nil {
			ErrorResponse(w, r, http.StatusUnauthorized, "Not Authorized")
			return
		}
		// Add user ID to request context
		ctx := context.WithValue(r.Context(), "userid", claims.UserID)
		next(w, r.WithContext(ctx))
	})
}