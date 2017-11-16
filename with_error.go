package smile

import (
	"errors"
	"math/rand"
)

func WithError() (string, error) {
	if rand.Intn(100)%2 == 0 {
		return "divisible by 2", nil
	}
	return "", errors.New("not divisible by 2")
}
