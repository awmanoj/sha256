package lib

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Module abstracts the SHA256 hash module
type SHA256Module struct{
	Salt string
}

// NewSHA256Module initializes the module with default non-salt hash.Hash
func NewSHA256Module() *SHA256Module {
	return &SHA256Module{}
}

// SHA256 will perform a SHA256 encryption and return the encrypted string
func (m *SHA256Module) SHA256(message string, saltKey string) (encrypted string) {
	hash := sha256.New()
	hash.Write([]byte(message + saltKey))
	encrypted = hex.EncodeToString(hash.Sum(nil))
	return encrypted
}

// DoubleSHA256 does a sha256(sha256(...)) operation
func (m *SHA256Module) DoubleSHA256(message string, saltKey string) (encrypted string) {
	return m.SHA256(m.SHA256(message, saltKey), "")
}