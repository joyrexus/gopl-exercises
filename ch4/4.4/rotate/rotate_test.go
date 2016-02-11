package rotate_test

import (
	"reflect"
	"testing"

	. "github.com/joyrexus/gopl-exercises/ch4/4.4/rotate"
)

func TestRotate(t *testing.T) {

	tests := []struct {
		by   int
		data []int
		want []int
	}{
		{3, []int{1, 2, 3, 4, 5}, []int{4, 5, 1, 2, 3}},
		{1, []int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}},
		{2, []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
		{5, []int{1, 2, 3}, []int{3, 1, 2}},
		{1, []int{}, []int{}},
	}

	for _, test := range tests {
		Rotate(test.by, test.data)
		if !reflect.DeepEqual(test.data, test.want) {
			t.Errorf("%v != %v", test.data, test.want)
		}
	}
}
