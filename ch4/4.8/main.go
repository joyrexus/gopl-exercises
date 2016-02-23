// See page 97.

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	counts := make(map[string]int)  // counts of Unicode character types
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		switch {
		case unicode.IsControl(r):
			counts["control"]++
		case unicode.IsLetter(r):
			counts["letter"]++
		case unicode.IsMark(r):
			counts["mark"]++
		case unicode.IsNumber(r):
			counts["number"]++
		case unicode.IsPunct(r):
			counts["punct"]++
		case unicode.IsSpace(r):
			counts["space"]++
		case unicode.IsSymbol(r):
			counts["symbol"]++
		}
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
