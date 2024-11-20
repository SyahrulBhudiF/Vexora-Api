package utils

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

func HashPassword(password string, salt string) string {
	salted := password + salt
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(salted), bcrypt.DefaultCost)
	hashed := string(hashedByte)

	if err != nil {
		panic(err)
	}

	return hashed
}

func ComparePassword(hashedPassword, password, salt string) bool {
	salted := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(salted))
	return err == nil
}

func GenerateOTP() string {
	mx := big.NewInt(1000000)

	n, _ := rand.Int(rand.Reader, mx)

	return fmt.Sprintf("%06d", n)
}
