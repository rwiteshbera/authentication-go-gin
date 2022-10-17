package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UserId    string
	jwt.RegisteredClaims
}

func GenerateToken(email string, firstname string, lastname string, uid string) (signedToken string, err error) {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal(envErr)
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")

	claims := SignedDetails{
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		UserId:    uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		fmt.Println("error ", err)
		return
	}

	return token, err
}