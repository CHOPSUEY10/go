package auth

import (
	"crypto/rand"
	"crypto/subtle"

	"golang.org/x/crypto/argon2"
)

const (
	saltLength = 16       // length of salt in bytes
	keyLength  = 32       // length of derived key
	time       = 1        // number of iterations
	memory     = 64 << 10 // mem cost in KiB (~64MB)
	threads    = 4
)

func generateSalt(len int) ([]byte, error) {

	salt := make([]byte, len)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

func HashPassword(password string) ([]byte, []byte, error) {
	salt, err := generateSalt(saltLength)
	if err != nil {
		return nil, nil, err
	}

	hashed := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)

	return hashed, salt, nil
}

func VerifyPassword(password string, salt []byte, expectedHash []byte) bool {

	newHash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)

	return subtleCompare(newHash, expectedHash)

}

func subtleCompare(a, b []byte) bool {

	return subtle.ConstantTimeCompare(a, b) == 1
}
