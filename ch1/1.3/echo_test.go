package echo

import (
	"strings"
	"testing"
)

var result string

func benchmarkAlpha(args []string, b *testing.B) {
	var r string
    for i := 0; i < b.N; i++ {
        r = Alpha(args)
    }
	result = r
}

func benchmarkBeta(args []string, b *testing.B) {
	var r string
    for i := 0; i < b.N; i++ {
        r = Beta(args)
    }
	result = r
}

func BenchmarkAlpha1(b *testing.B) {
	benchmarkAlpha([]string{"x"}, b)
}

func BenchmarkAlpha2(b *testing.B) {
	args := []string{"a", "b", "c", "d"}
	benchmarkAlpha(args, b)
}

func BenchmarkAlpha3(b *testing.B) {
	args := strings.Split("abcdefghijk", "")
	benchmarkAlpha(args, b)
}

func BenchmarkBeta1(b *testing.B) {
	benchmarkBeta([]string{"x"}, b)
}

func BenchmarkBeta2(b *testing.B) {
	args := []string{"a", "b", "c", "d"}
	benchmarkBeta(args, b)
}

func BenchmarkBeta3(b *testing.B) {
	args := strings.Split("abcdefghijk", "")
	benchmarkBeta(args, b)
}
