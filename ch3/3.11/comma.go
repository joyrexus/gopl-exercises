package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("123"))         // 123
	fmt.Println(comma("1234"))        // 1,234
	fmt.Println(comma("12345"))       // 12,234
	fmt.Println(comma("123456"))      // 123,456
	fmt.Println(comma("1234567"))     // 1,234,567
	fmt.Println(comma("1234567.01"))  // 1,234,567.01
	fmt.Println(comma("-1234567.01")) // -1,234,567.01
}

// comma inserts commas in a decimal integer or float string.
func comma(s string) string {
	var buf bytes.Buffer

	// strip sign if present
	if strings.HasPrefix(s, "-") {
		buf.WriteString("-")
		s = strings.TrimPrefix(s, "-")
	}

	// extract fractional part of floats
	var frac string
	if strings.Contains(s, ".") {
		parts := strings.SplitN(s, ".", 2)
		s, frac = parts[0], parts[1]
	}

	// no commas needed
	if len(s) <= 3 {
		if frac != "" {
			s += "." + frac
		}
		return s
	}

	l := len(s)
	r := l % 3 // remainder modulo 3 to get initial digits

	// insert comma after initial digits
	if r >= 1 {
		buf.WriteString(s[:r])
		buf.WriteString(",")
	}
	// insert commas after each remaining set of 3 digits
	for i := r; i < l; i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < l {
			buf.WriteString(",")
		}
	}

	// add fractional remainder
	if frac != "" {
		buf.WriteString(".")
		buf.WriteString(frac)
	}
	return buf.String()
}
