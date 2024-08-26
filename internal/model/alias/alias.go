package alias

import "golang.org/x/exp/rand"

const (
	keys      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	keysCount = len(keys)
)

var (
	runeKeys = []rune(keys)
)

func New(length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = runeKeys[rand.Intn(keysCount)]
	}

	return string(result)
}
