package crypto_wallet_generator

import (
	"math/rand"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateRandomString(length uint) string {

	res := make([]byte, length)

	for i := range letters {
		res[i] = letters[rand.Intn(len(letters))]
	}

	return string(res)
}
