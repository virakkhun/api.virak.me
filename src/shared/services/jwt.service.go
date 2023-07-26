package sharedServices

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"api.virak.me/src/shared/models"
	"api.virak.me/src/utils"
	"github.com/golang-jwt/jwt/v5"
)

func SignNewToken(args models.Map) (string, error) {
	fmt.Println(args)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": args["ID"],
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	secret := []byte(os.Getenv("SECRET"))

	tokenString, err := token.SignedString(secret)

	if utils.IsNotNil(err) {
		return tokenString, errors.New("failed to generated tokenString")
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, error) {
	secret := []byte(os.Getenv("SECRET"))

	verified, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if !verified.Valid {
		return verified.Valid, errors.New(strings.Join([]string{`Unable to verified tokenString`, err.Error()}, ", "))
	}

	return verified.Valid, nil
}
