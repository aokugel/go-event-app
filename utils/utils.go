package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var key = []byte(os.Getenv("TOKEN_SECRET"))

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}

func GetToken(id int64, email string, name string) (signedString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserId": id,
			"iss":    "anthonykugel.com",
			"sub":    email,
			"name":   name,
			"exp":    time.Now().Add(time.Hour * 2).Unix(),
		})
	signedString, err = token.SignedString(key)

	if err != nil {
		return "", err
	}

	return signedString, nil

}
