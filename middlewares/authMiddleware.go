package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// Middleware function to validate JWT tokens
func JWTAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        err := godotenv.Load()
         if err != nil {
        log.Fatal("Error loading .env file")
    }
        // Get the token from the Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header missing", http.StatusUnauthorized)
            return
        }

        // Check if the token is prefixed with "Bearer "
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
            return
        }

        tokenString := parts[1]

        // Parse and validate the token
        token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("SECRET_KEY")), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Extract claims and add to the context if needed
        if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
            ctx := context.WithValue(r.Context(), "user_id", claims.Subject)
            next.ServeHTTP(w, r.WithContext(ctx))
        } else {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
        }
    })
}
