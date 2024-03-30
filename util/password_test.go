package util

import "testing"

func TestPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Error(err)
	}

	err = CheckPasswordHash(password, hash)
	if err != nil {
		t.Error(err)
	}
}
