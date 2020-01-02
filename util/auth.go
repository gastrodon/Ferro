package util

var secret string

func RecieveSecret(recieved string) {
	secret = recieved
}

func Authed(auth string) (authed bool) {
	return auth == secret
}
