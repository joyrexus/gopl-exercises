// Dup3 prints the count and text of lines that appear more than once
// in the input.  It takes a list of named files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)
	for _, fname := range os.Args[1:] {
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "duplicates: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			if line != "" {
				count[line]++
			}
		}
	}
	
	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
