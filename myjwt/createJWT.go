package myjwt

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("my-secret")

func CreateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(5 * time.Minute).Unix()

	tokenString, err := token.SignedString(SECRET)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unauthorized")
		}
		return SECRET, nil
	})

	if err != nil {
		return err
	}
	if token.Valid {
		return nil
	}
	return errors.New("Unauthorized")
}
