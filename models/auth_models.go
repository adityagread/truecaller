package models

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load .env file")
	}
}

var JwtSecret = []byte(os.Getenv("JWT_KEY"))

type CustomClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}
