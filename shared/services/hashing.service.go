package shared_services

import (
	"log"

	"api.virak.me/utils"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if utils.IsNotNil(err) {
		log.Fatalln("failed to generate hash password")
	}

	return string(hashedPassword)
}

// return true if compare success, false if failed
func DeHashPassword(hashedPass string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))

	return !utils.IsNotNil(err)
}
