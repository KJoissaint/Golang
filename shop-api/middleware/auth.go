package middleware

import (
	"context"
	"net/http"
	"shop-api/models"
	"shop-api/utils"
	"strings"
)

type contextKey string

const ClaimsContextKey contextKey = "claims"

// AuthMiddleware validates JWT token and adds claims to context
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "authorization header required"}`, http.StatusUnauthorized)
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, `{"error": "invalid authorization header format"}`, http.StatusUnauthorized)
			return
		}

		token := parts[1]
		claims, err := utils.ValidateToken(token)
		if err != nil {
			http.Error(w, `{"error": "invalid or expired token"}`, http.StatusUnauthorized)
			return
		}

		// Add claims to context
		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
		next(w, r.WithContext(ctx))
	}
}

// RequireSuperAdmin middleware ensures the user is a SuperAdmin
func RequireSuperAdmin(next http.HandlerFunc) http.HandlerFunc {
	return AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(ClaimsContextKey).(*utils.Claims)
		if !ok {
			http.Error(w, `{"error": "invalid context"}`, http.StatusInternalServerError)
			return
		}

		if claims.Role != models.RoleSuperAdmin {
			http.Error(w, `{"error": "super admin access required"}`, http.StatusForbidden)
			return
		}

		next(w, r)
	})
}

// RequireAdmin middleware ensures the user is at least an Admin
func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(ClaimsContextKey).(*utils.Claims)
		if !ok {
			http.Error(w, `{"error": "invalid context"}`, http.StatusInternalServerError)
			return
		}

		if claims.Role != models.RoleSuperAdmin && claims.Role != models.RoleAdmin {
			http.Error(w, `{"error": "admin access required"}`, http.StatusForbidden)
			return
		}

		next(w, r)
	})
}

// GetClaims extracts claims from context
func GetClaims(r *http.Request) (*utils.Claims, bool) {
	claims, ok := r.Context().Value(ClaimsContextKey).(*utils.Claims)
	return claims, ok
}
