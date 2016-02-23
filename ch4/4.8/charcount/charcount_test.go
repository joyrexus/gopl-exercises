package charcount_test

import (
	"bufio"
	"strings"
	"testing"

	. "github.com/joyrexus/gopl-exercises/ch4/4.8/charcount"
)

var tests = []struct {
	input string
	want  map[string]int
}{
	{"", map[string]int{}},
	{"世界", map[string]int{"letter": 2}},
	{"\x90", map[string]int{"invalid": 1}},
	{"hello, world!", map[string]int{
		"space":  1,
		"letter": 10,
		"punct":  2,
	}},
	{"\tWAT!\n", map[string]int{
		"control": 2,
		"letter":  3,
		"punct":   1,
	}},
	{"a + b = c", map[string]int{
		"symbol": 2,
		"letter": 3,
		"space":  4,
	}},
}

func TestCharCount(t *testing.T) {
	for _, test := range tests {
		r := strings.NewReader(test.input)
		counts := CharCount(bufio.NewReader(r))
		if !eq(counts, test.want) {
			t.Errorf("got %v, want %v", counts, test.want)
		}
	}
}

// eq checks whether two map[string]ints are equal.
func eq(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
