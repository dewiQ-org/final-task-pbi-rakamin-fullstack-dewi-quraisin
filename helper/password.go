package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), err
}

func VerifyPassword(hashPasssword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPasssword), []byte(password))
	return err
}
