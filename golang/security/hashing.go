package security

import "golang.org/x/crypto/bcrypt"

const SaltRounds = 10

// HashPassword hashes a password using bcrypt with a salt length.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), SaltRounds)
	return string(bytes), err
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent. Returns true if the password is correct.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
