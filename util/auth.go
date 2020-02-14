package util

import (
	"os"
)

var (
	secret = os.Getenv("FERRO_SECRET")
)

func Authed(auth string) (authed bool) {
	return auth == secret
}
