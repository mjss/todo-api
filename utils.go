package main

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func verifyPassword(storedHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))

	if err != nil {
		return false
	}

	return true
}
