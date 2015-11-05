package echo

import (
	"strings"
)

// original echo implementation
func Alpha(args []string) string {
	var s, sep string
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

// a more efficient echo implementation?
func Beta(args []string) string {
	return strings.Join(args[1:], " ")
}
