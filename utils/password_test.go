package utils_test

import (
	"API_BIG_WORK/utils"
	"log"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "xyh666"
	salt := "6666"

	hashedPassword, err := utils.HashPasswordWithSalt(password, salt)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	log.Print(hashedPassword)
}

func TestComparePasswords(t *testing.T) {
	password := "xyh666"
	salt := "6666"

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if !utils.ComparePasswordsWithSalt(hashedPassword, password, salt) {
		t.Errorf("Passwords do not match")
	}
}
