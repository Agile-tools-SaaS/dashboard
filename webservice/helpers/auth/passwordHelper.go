package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(hashedPassword, entered_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(entered_password))
	return err
}
