package routes

import (
	"context"
	"main_server/utils"
	"net/http"
)

type MiddlewareBody struct {
	Protected   bool
	Handlerfunc http.Handler
}

func Protected(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(w, r, http.StatusUnauthorized, "Authorization header missing")
			return
		}

		// Ensure it's a Bearer token
		const bearerPrefix = "Bearer "
		if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			utils.ErrorResponse(w, r, http.StatusUnauthorized, "Invalid token format")
			return
		}

		// Extract the actual token
		token := authHeader[len(bearerPrefix):]

		claims, err := utils.DecodeToken(token)
		if err != nil {
			utils.ErrorResponse(w, r, http.StatusUnauthorized, "Not Authorized")
			return
		}
		// Add user ID to request context
		ctx := context.WithValue(r.Context(), "userid", claims.UserID)
		next(w, r.WithContext(ctx))
	})
}
