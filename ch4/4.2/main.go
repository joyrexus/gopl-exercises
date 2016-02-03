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

func main() {
	s384 := flag.Bool("stronger", false, "Print SHA384 hash")
	s512 := flag.Bool("strongest", false, "Print SHA512 hash")
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Could not read STDIN:", err)
	}

	var hash string	// digest of message on stdin in hexadecimal

	switch {
	case *s384:
		hash = fmt.Sprintf("%x", sha512.Sum384(bytes))
	case *s512:
		hash = fmt.Sprintf("%x", sha512.Sum512(bytes))
	default:
		hash = fmt.Sprintf("%x", sha256.Sum256(bytes))
	}
	fmt.Println(hash)
}
