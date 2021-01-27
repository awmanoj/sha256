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

	fmt.Println(lib.NewSHA256Module().EncodeWithSalt(data, secret))
	fmt.Println(lib.NewSHA256Module().Encode(data))
}
