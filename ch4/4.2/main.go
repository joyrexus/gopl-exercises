package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// print digest of stdin in hexadecimal
func main() {
	sha := flag.String("sha", "256", "256 | 384 | 512")
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Could not read STDIN:", err)
	}

	switch *sha {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256(bytes))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384(bytes))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512(bytes))
	default:
		fmt.Printf("%s is not a valid option: 256 | 383 | 512\n", *sha)
	}
}
