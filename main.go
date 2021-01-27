package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Printf("Usage: %s secret data\n", os.Args[0])
		return
	}

	secret := os.Args[1]
	data := os.Args[2]

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha256 := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sha256)
}
