package rev_test

import (
	"testing"

	"github.com/joyrexus/gopl-exercises/ch4/4.3/rev"
)

func TestReverse(t *testing.T) {

	tests := []struct {
		input [6]int
		want  [6]int
	}{
		{
			[...]int{0, 1, 2, 3, 4, 5},
			[...]int{5, 4, 3, 2, 1, 0},
		},
		{
			[...]int{5, 4, 3, 2, 1, 0},
			[...]int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		if rev.Reverse(&test.input); test.input != test.want {
			t.Errorf("%v != %v", test.input, test.want)
		}
	}
}
