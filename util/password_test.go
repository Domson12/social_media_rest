package util

import "testing"

func TestPassword(t *testing.T) {
	t.Run("HashPassword", func(t *testing.T) {
		password := "password"
		hashedPassword, err := HashPassword(password)

		if err != nil {
			t.Error(err)
		}

		if hashedPassword == password {
			t.Error("Hashed password is the same as the original password")
		}
	})

	t.Run("CheckPasswordHash", func(t *testing.T) {
		password := "password"
		hashedPassword, err := HashPassword(password)

		if err != nil {
			t.Error(err)
		}

		err = CheckPasswordHash(password, hashedPassword)

		if err != nil {
			t.Error("Password and hash do not match")
		}
	})
}
