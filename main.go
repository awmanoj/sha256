package main

import (
	"fmt"
	"github.com/awmanoj/sha256/lib"
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

	fmt.Printf("sha256, with salt, '%s':\n%s\n", secret, lib.NewSHA256Module().SHA256(data, secret))
	fmt.Printf("double-sha256, with salt, '%s':\n%s\n", secret, lib.NewSHA256Module().DoubleSHA256(data, secret))
}
