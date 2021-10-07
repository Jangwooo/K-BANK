package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Hash(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(hash)

}
