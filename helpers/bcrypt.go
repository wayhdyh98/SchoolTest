package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(text string) string {
	password := []byte(text)
	hash, _ := bcrypt.GenerateFromPassword(password, 8)

	return string(hash)
}

func ComparePass(hashText, passText []byte) bool {
	hash, pass := []byte(hashText), []byte(passText)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}