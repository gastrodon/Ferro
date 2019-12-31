package util

import (
	"errors"
	"fmt"
)

func TooManyResults(size int) (err error) {
	err = errors.New(fmt.Sprintf("Too many results: %d", size))
	return
}
