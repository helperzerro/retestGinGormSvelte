package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte("your-secret-key")

func GenerateJWT(email, role string) (string, error) {
	// Durasi token: 26 jam
	// expirationTime := time.Now().Add(26 * time.Hour)
	expirationTime := time.Now().Add(3 * time.Second)


	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   expirationTime.Unix(), // UNIX timestamp
		"iat":   time.Now().Unix(),             // issued at
    "jti":   uuid.NewString(),              // unique token ID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tanda tangan token pakai secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
