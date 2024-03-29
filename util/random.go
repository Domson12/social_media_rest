package util

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := 0; i < 10; i++ {
		randInt := r.Intn(100)
		println(randInt)
	}
}

func RandomInt(min, max int64) int {
	return int(min + rand.Int63n(max-min+1))
}

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(letterBytes[rand.Intn(len(letterBytes))])
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}
