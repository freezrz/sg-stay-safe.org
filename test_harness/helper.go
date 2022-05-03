package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomInt64() int64 {
	i := rand.Int63n(int64(9999999999))
	r := i + 10000000000
	return r
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
