package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword : Takes in a string and a salt and generates a hash
func HashPassword(password string, salt int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

// CheckPasswordHash : Compares the password with the password hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
