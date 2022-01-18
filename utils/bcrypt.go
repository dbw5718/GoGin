package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Panicln(err)
	}
	return string(hash)
}
func BcryMakeCheck(pwd []byte, hashPwd string) bool {
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		return false
	}
	return true
}
