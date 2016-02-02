package shadiff_test

import (
	"crypto/sha256"
	"testing"

	"github.com/joyrexus/gopl-exercises/ch4/4.1/shadiff"
)

var x = sha256.Sum256([]byte("x"))
var y = sha256.Sum256([]byte("X"))

func BenchmarkCountByBit(b *testing.B) {
    for i := 0; i < b.N; i++ {
        shadiff.CountByBit(x, y)
    }
}

func BenchmarkCountByByte(b *testing.B) {
    for i := 0; i < b.N; i++ {
        shadiff.CountByByte(x, y)
    }
}

// BenchmarkCountByBit-4 	  200000	      8496 ns/op
// BenchmarkCountByByte-4	 5000000	       353 ns/op
