package util

import (
	"os"
)

var (
	secret = os.Getenv("MONKE_SECRET")
)

func Authed(auth string) (authed bool) {
	return auth == secret
}
