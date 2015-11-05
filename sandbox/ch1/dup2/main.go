// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, count)
		displayCounts(count)
		return
	}
	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "duplicates: %v\n", err)
			continue
		}
		countLines(f, count)
		f.Close()
	}
	displayCounts(count)
}
	
func displayCounts(count map[string]int) {
	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}
