package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type SHA256Module struct{}

func NewSHA256Module() *SHA256Module {
	return &SHA256Module{}
}

func (*SHA256Module) EncodeWithSalt(message string, saltKey string) (encrypted string) {
	h := hmac.New(sha256.New, []byte(saltKey))
	h.Write([]byte(message))
	encrypted = hex.EncodeToString(h.Sum(nil))
	return encrypted
}

func (*SHA256Module) Encode(message string) (encrypted string) {
	h := sha256.New()
	h.Write([]byte(message))
	encrypted = hex.EncodeToString(h.Sum(nil))
	return encrypted
}
