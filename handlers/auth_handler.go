package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/gupta799/go_api/response"
	"github.com/joho/godotenv"
)

type MyCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type JwtToken struct {
	JwtToken string `json:"token"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    secret_key:=os.Getenv("SECRET_KEY")
	claims := MyCustomClaims{
		uuid.New().String(), // User ID
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 24 hours expiration
			Issuer:    "myapp",
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		fmt.Println("Error signing token:", err)
		response.RespondWithJson(w, 500, &JwtToken{
			tokenString,
		})

	}

	response.RespondWithJson(w, 200, &JwtToken{
		tokenString,
	})

}
