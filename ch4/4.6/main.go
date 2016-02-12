package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// squash squashes each run of adjacent Unicode spaces in a UTF-8 
// encoded []byte slice into a single ASCII space. It is an "in-place"
// function (see 4.22, p. 91) in that it modifies elements of the
// slice "in-place".
func squash(input []byte) []byte {
	i := 0
	runes := string(input)
	lastWasSpace := false
	for _, r := range runes {
		if unicode.IsSpace(r) || r == ' ' {
			if lastWasSpace {
				continue
			} else {
				lastWasSpace = true
				r = ' '
			}
		} else {
			lastWasSpace = false
		}
		if utf8.RuneLen(r) > 1 {
			buf := make([]byte, 3)
			n := utf8.EncodeRune(buf, r)
			j := i + n
			copy(input[i:j], buf)
			i = j
		} else {
			input[i] = byte(r)
			i++
		}
	}
	return input[:i]
}

// trim is a variant of squash, trimming spaces.
func trim(input []byte) []byte {
	fields := strings.Fields(string(input))
	s := strings.Join(fields, " ")
	return []byte(s)
}

func main() {
	s := []byte("x\t\t\t世\t\t\ty \v\v\n  z")
	// s := []byte("x\t\t\tx\t\t\tx    x")
	s = trim(s)
	fmt.Printf("%s", s) // "x 世 y z"
}
