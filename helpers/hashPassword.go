package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPwd(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
