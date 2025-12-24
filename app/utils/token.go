package utils

import (
	"os"
	"strconv"
	"time"
)

type TokenPayload struct {
	Username string    `json:"username"`
	Expired  time.Time `json:"expired"`
	Type     string    `json:"type"`
}

func GetToken(username string) (string, error) {
	expMin, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_MINUTES"))

	payload := TokenPayload{
		Username: username,
		Expired:  time.Now().Add(time.Minute * time.Duration(expMin)),
		Type:     "access",
	}

	return Encrypt(payload)
}
