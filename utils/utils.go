package utils

import (
	"errors"
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
			"userId": id,
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

func ParseToken(accessToken string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("token isn't correctly signed")
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claim is of incorrect type")
	}

	return claims, nil
}
