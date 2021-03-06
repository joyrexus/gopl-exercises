package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte // 256 = 8^2 possible combinations of 0 and 1 for 8 slots

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var s byte
	for i := uint64(0); i < 8; i++ {
		s += pc[byte(x>>(i*8))]
	}
	return int(s)
}

func main() {
	for _, x := range []uint64{0, 1, 2, 3, 255, 0x1234567890ABCDEF} {
		fmt.Printf("\n% 20d has a popcount of %d: %08[1]b\n", x, PopCount(x))
	}
}
