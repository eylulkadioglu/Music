package utils

import (
	"time"

	"github.com/eylulkadioglu/Music/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetCode() string {
	uuid := uuid.New()

	return uuid.String()
}

// Returns the private key that is used for signing JWT tokens
// This key is used for token verification and signing operations
func GetJwtKey() []byte {
	return []byte("eylul123888")
}


// Generates a JWT token that includes the given user data 
func GetJwtToken(user models.User) (string, error) {
	expiration := time.Now().Add(24 * time.Hour) // Limiting the validity period with 24 hours

	// Generates the required claims for token
	claims := &models.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	// Creates a new token with HS256 algorithm
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	// Signs the token and returns as a string
	tokenString, err := token.SignedString(GetJwtKey())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
