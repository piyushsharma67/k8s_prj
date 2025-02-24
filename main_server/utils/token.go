package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecretKey = []byte("dgdsgdfgdfsg23452342342342") // Use a secure key in production

// Claims struct represents the custom claims you want to include in the JWT
type Claims struct {
	UserID int32 `json:"user_id"`
	jwt.StandardClaims
}

// EncodeToken generates a JWT token for a given userID
func EncodeToken(userID int32) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	// Create the claims (the payload of the token)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Set expiration
			Issuer:    "my-app",              // Optional: Issuer field
		},
	}

	// Create the token using the claims and sign it with a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// DecodeToken decodes and validates the JWT token, extracting the claims
func DecodeToken(tokenStr string) (*Claims, error) {
	// Parse the token and validate it using the secret key
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract the claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
