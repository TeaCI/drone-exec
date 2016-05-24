package token

import "github.com/dgrijalva/jwt-go"

// New generates a new agent token for authenticating to the drone server.
func New(secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["type"] = "agent"
	token.Claims["text"] = ""
	return token.SignedString([]byte(secret))
}
