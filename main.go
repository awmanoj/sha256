package main

import (
	"fmt"
	"github.com/awmanoj/sha256/lib"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Printf("Usage: %s message salt \n", os.Args[0])
		return
	}

	message := os.Args[1]
	salt := os.Args[2]

	m := lib.NewSHA256Module()

	fmt.Printf("%s%s\n", message, salt)
	fmt.Printf("SHA256 = %s\n", m.SHA256(message, salt))
	fmt.Printf("DoubleSHA256 = %s\n", m.DoubleSHA256(message, salt))
}
