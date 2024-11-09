package passwordLib

import "golang.org/x/crypto/bcrypt"

// Validate a password
func Validate(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
