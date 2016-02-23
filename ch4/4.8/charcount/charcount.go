// See page 97.

// charcount computes counts of Unicode characters.
package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// CharCount scans a *bufio.Reader and returns a map of the counts of its
// Unicode character types.
func CharCount(in *bufio.Reader) map[string]int {
	counts := make(map[string]int) // counts of Unicode character types

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		switch {
		case r == unicode.ReplacementChar && n == 1:
			counts["invalid"]++
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
	}
	return counts
}
