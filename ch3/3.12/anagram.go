package main

import (
	"fmt"
)

func main() {
	a := "strings"
	b := "grinsts"
	fmt.Println(isAnagram(a, b)) // true

	a = "baz"
	b = "zab"
	fmt.Println(isAnagram(a, b)) // true

	a = "bar"
	b = "baz"
	fmt.Println(isAnagram(a, b)) // false
}

// isAnagram tests whether two strings are anagrams of one another.
func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aMap := make(map[rune]int) // count of each rune in a
	for _, r := range a {
		aMap[r]++
	}

	// Ensure that each rune r in b is also in a and that
	// the current count or r in b is < the count in a.
	bMap := make(map[rune]int)
	for _, r := range b {
		bMap[r]++
		if n, ok := aMap[r]; !ok || n < bMap[r] {
			return false
		}
	}

	return true
}
