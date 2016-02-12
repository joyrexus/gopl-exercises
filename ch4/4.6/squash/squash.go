package squash

import (
	"unicode"
	"unicode/utf8"
	"strings"
)

// Squash squashes each run of adjacent Unicode spaces in a UTF-8 
// encoded []byte slice into a single ASCII space. It is an "in-place"
// function (see 4.22, p. 91) in that it modifies elements of the
// slice "in-place".
func Squash(input []byte) []byte {
	i := 0
	runes := string(input)
	prevSpace := false // was the previous rune a space?
	for _, r := range runes {
		if unicode.IsSpace(r) || r == ' ' {
			if prevSpace {
				continue
			} else {
				prevSpace = true
				r = ' '	// convert to an ascii space
			}
		} else {
			prevSpace = false
		}
		if utf8.RuneLen(r) > 1 {
			buf := make([]byte, 3)
			j := i + utf8.EncodeRune(buf, r)
			copy(input[i:j], buf)
			i = j
		} else {
			input[i] = byte(r)
			i++
		}
	}
	return input[:i]
}

// Trim is a variant of squash, trimming adjacent spaces.
func Trim(input []byte) []byte {
	fields := strings.Fields(string(input))
	s := strings.Join(fields, " ")
	return []byte(s)
}

