package main

import "fmt"

func rotate(by int, s []int, right bool) []int {
	if by == 0 || len(s) == 0 {
		return s
	}
	i := by % len(s)
	j := len(s) - i
	if right {
		i, j = j, i
	}
	q := make([]int, len(s))
	copy(q, s[i:])
	copy(q[j:], s[:i])
	return q
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5} // array
	// Rotate s right by two positions.
	s = rotate(8, s, true) // "[4 5 0 1 2 3]"
	fmt.Println(s)
	s = rotate(2, s, false) // "[0 1 2 3 4 5]"
	fmt.Println(s)
}
