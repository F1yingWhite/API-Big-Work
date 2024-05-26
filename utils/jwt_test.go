package utils_test

import (
	"API_BIG_WORK/utils"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := utils.CreateToken("xyh")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	token, err := utils.CreateToken("xyh")
	if err != nil {
		t.Error(err)
	}
	jwt, err := utils.ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(jwt.ID)
}
