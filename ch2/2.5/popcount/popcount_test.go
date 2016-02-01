package popcount_test

import (
	"testing"

	pc "github.com/joyrexus/gopl-exercises/ch2/2.5/popcount"
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

func BenchmarkPopcountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.PopCountShift(0x1234567890ABCDEF)
	}
}

func BenchmarkPopcountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.PopCountClear(0x1234567890ABCDEF)
	}
}

// BenchmarkPopcount-4    	200000000	         8.43 ns/op
// BenchmarkPopcountLoop-4	100000000	        22.8 ns/op
// BenchmarkPopcountShift-4	10000000	       132 ns/op
// BenchmarkPopcountClear-4	50000000	        32.9 ns/op

