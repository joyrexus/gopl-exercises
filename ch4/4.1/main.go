package main

import (
	"crypto/sha256"
	"fmt"

	pc "gopl.io/ch2/popcount"
)

func CountByBit(a, b [32]byte) int {
    n := 0
	for i := range a {
		for j := uint(0); j < 64; j++ {
			if a[i]&(1<<j) != b[i]&(1<<j) {
				n++
			}
		}
		fmt.Printf("\n%08b\n%08b\n%d\n\n", a[i], b[i], n)
	}
    return n
}

func CountByByte(a, b [32]byte) int {
	n := 0
	for i := range a {
		n += pc.PopCount(uint64(a[i] ^ b[i]))
	}
	return n
}

func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))
	fmt.Println(diffcount(a, b))
	fmt.Println(diff(a, b))
}
