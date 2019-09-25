package cats

import (
	"errors"
	"strings"
)

func Meow(repeat int) (string, error) {
	if repeat < 0 {
		return "", errors.New("not enough meows")
	}
	return strings.Repeat("meow", repeat), nil
}
