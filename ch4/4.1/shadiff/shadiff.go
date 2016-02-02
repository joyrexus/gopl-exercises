package shadiff

import (
	pc "gopl.io/ch2/popcount"
)

// CountByBit returns the number of bits that differ between
// two SHA256 hashes. It compares each bit in the corresponding
// bytes of each hash.
func CountByBit(a, b [32]byte) int {
    n := 0
	for i := range a {
		for j := uint(0); j < 64; j++ {
			if a[i]&(1<<j) != b[i]&(1<<j) {
				n++
			}
		}
	}
    return n
}

// CountByByte returns the number of bits that differ between
// two SHA256 hashes. It sums the population count of the bitwise 
// difference between corresponding bytes in each hash.
func CountByByte(a, b [32]byte) int {
	n := 0
	for i := range a {
		n += pc.PopCount(uint64(a[i]^b[i]))
	}
	return n
}


// BenchmarkCountByBit-4 	  200000	      8739 ns/op
// BenchmarkCountByByte-4	 5000000	       343 ns/op
