package main

import "fmt"

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}

func main() {
	for _, x := range []uint64{0, 1, 2, 3, 255, 0x1234567890ABCDEF} {
		fmt.Printf("\n% 20d has a popcount of %d: %08[1]b\n", x, PopCount(x))
	}
}
