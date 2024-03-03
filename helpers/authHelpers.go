package helpers

import (
	"aditya-coding-task/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(models.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
