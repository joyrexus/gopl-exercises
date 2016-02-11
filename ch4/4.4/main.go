package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5} // array
	// Rotate s right by two positions.
	q := make([]int, len(s))
	
	by := 2
	i := len(s) - by
	copy(q, s[i:])
	copy(q[by:], s[:i])
	copy(s, q)
	fmt.Println(s) // "[4 5 0 1 2 3]"
}
