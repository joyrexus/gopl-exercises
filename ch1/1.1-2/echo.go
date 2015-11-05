package main

import (
	"fmt"
	"os"
)

func main() {
	for i, x := range os.Args[:] {
		fmt.Println(i, x)
	}
}
