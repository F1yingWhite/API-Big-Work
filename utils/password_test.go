package utils_test

import (
	"API_BIG_WORK/utils"
	"log"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "5050"

	hashedPassword, err := utils.HashPassword(password, "afXHSqu0IS1kEDod3pdE")
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	log.Print(hashedPassword)
}

func TestComparePasswords(t *testing.T) {
	password := "5050"
	salt := "afXHSqu0IS1kEDod3pdE"

	hashedPassword, err := utils.HashPassword(password, salt)
	log.Print(hashedPassword)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if !utils.ComparePasswords(hashedPassword, password, salt) {
		t.Errorf("Passwords do not match")
	}
	print("对的对的")
}
