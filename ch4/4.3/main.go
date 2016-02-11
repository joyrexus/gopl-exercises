package main

import "fmt"

// reverse reverses an array of 6 ints in place.
func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5} // array
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
}
