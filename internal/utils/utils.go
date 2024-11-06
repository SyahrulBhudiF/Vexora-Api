package utils

import "golang.org/x/crypto/bcrypt"

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
