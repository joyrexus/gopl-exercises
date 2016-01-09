package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123"))     // 123
	fmt.Println(comma("1234"))    // 1,234
	fmt.Println(comma("12345"))   // 12,345
	fmt.Println(comma("123456"))  // 123,456
	fmt.Println(comma("1234567")) // 1,234,567
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	// no commas needed
	if len(s) <= 3 {
		return s
	}

	var buf bytes.Buffer
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

	return buf.String()
}
