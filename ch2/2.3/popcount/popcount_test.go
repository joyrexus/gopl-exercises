package popcount_test

import (
	"testing"

	pc "github.com/joyrexus/gopl-exercises/ch2/2.3/popcount"
)

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopcountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.PopCountLoop(0x1234567890ABCDEF)
	}
}


// BenchmarkPopcount-4    	200000000	         8.43 ns/op
// BenchmarkPopcountLoop-4	100000000	        22.8 ns/op
