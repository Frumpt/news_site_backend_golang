package users

import (
	"crypto/sha256"
	"encoding/hex"
)

func PasswordHasher(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)
}
