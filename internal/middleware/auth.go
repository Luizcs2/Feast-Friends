// Package middleware contains HTTP middleware functions for the application.
package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"feast-friends-api/internal/utils" // Assumes your utils package is in pkg/utils
	"net/http"
)

// contextKey is a custom type for our context key. It's a Go best practice
// to prevent collisions with other packages' context keys.
type contextKey string

// UserIDKey is the key we'll use to store and retrieve the user ID in the request context.
const UserIDKey contextKey = "userID"

// AuthMiddleware protects routes by verifying the JWT token from the Authorization header.
func AuthMiddleware(next http.Handler) http.Handler {
	// http.HandlerFunc is an adapter that allows the use of ordinary functions as HTTP handlers.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Attempt to retrieve the Authorization header from the incoming request.
		authHeader := r.Header.Get("Authorization")

		// If the header is missing, the request cannot be authenticated.
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			// Format a standard JSON error response.
			response := utils.ErrorResponse("Authorization header missing", errors.New("missing authorization header"), http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return // Terminate the request handling here.
		}

		// This function handles token verification against the auth provider.
		userID, err := utils.ValidateToken(authHeader)

		// If the token is invalid (e.g., expired, wrong signature), an error will be returned.
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			response := utils.ErrorResponse("Invalid or expired token", err, http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return // Terminate the request handling here.
		}

		// At this point, the user is authenticated.
		// We enrich the request's context with the validated userID.
		// This makes the userID available to subsequent handlers in the chain.
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		// The request is valid, so we pass it along to the next handler,
		// complete with the updated context.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}