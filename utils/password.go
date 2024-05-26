package utils

import (
	"API_BIG_WORK/config"
	"crypto/subtle"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := config.CFG.Salt
	time := uint32(1)
	memory := uint32(1024)
	threads := uint8(2)
	keyLen := uint32(32)

	hashedPassword := argon2.IDKey([]byte(password), []byte(salt), time, memory, threads, keyLen)

	return base64.StdEncoding.EncodeToString(hashedPassword), nil
}

func HashPasswordWithSalt(password string, salt string) (string, error) {
	time := uint32(1)
	memory := uint32(1024)
	threads := uint8(2)
	keyLen := uint32(32)

	hashedPassword := argon2.IDKey([]byte(password), []byte(salt), time, memory, threads, keyLen)

	return base64.StdEncoding.EncodeToString(hashedPassword), nil
}

func ComparePasswords(hashedPassword string, password string) bool {
	salt := config.CFG.Salt
	time := uint32(1)
	memory := uint32(1024)
	threads := uint8(2)
	keyLen := uint32(32)

	decodedHashedPassword, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		fmt.Println("Error decoding string", err.Error())
		return false
	}

	derivedKey := argon2.IDKey([]byte(password), []byte(salt), time, memory, threads, keyLen)

	return subtle.ConstantTimeCompare(decodedHashedPassword, derivedKey) == 1
}

func ComparePasswordsWithSalt(hashedPassword string, password string, salt string) bool {
	time := uint32(1)
	memory := uint32(1024)
	threads := uint8(2)
	keyLen := uint32(32)

	decodedHashedPassword, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		fmt.Println("Error decoding string", err.Error())
		return false
	}

	derivedKey := argon2.IDKey([]byte(password), []byte(salt), time, memory, threads, keyLen)

	return subtle.ConstantTimeCompare(decodedHashedPassword, derivedKey) == 1
}
