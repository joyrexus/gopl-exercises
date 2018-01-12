// Dup takes a list of filenames and reports the number of times a line of text
// occurs if it happens to occur more than once.  It also indicates per file
// counts for each duplicate found.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// For each line, keep track of total instances and per file counts
	type Tally struct {
		Files map[string]int	// per file counts
		Total int
	}
	count := make(map[string]*Tally)
	files := os.Args[1:]

	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "duplicates: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			line := input.Text()
			t, ok := count[line]
			if !ok {
				t = &Tally{
					Files: make(map[string]int),
					Total: 0,
				}
				count[line] = t
			}
			t.Files[fname]++
			t.Total += 1
		}
		f.Close()
	}

	for line, tally := range count {
		if tally.Total > 1 {
			fmt.Printf("%q occurred %d times\n", line, tally.Total)
			for file, count := range tally.Files {
				fmt.Printf("  %s: %d times\n", file, count)
			}
		}
	}
}
