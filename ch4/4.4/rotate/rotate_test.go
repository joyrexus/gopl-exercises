package rotate_test

import (
	"reflect"
	"testing"

	. "github.com/joyrexus/gopl-exercises/ch4/4.4/rotate"
)

func TestRotate(t *testing.T) {

	tests := []struct {
		by    int
		right bool
		data  []int
		want  []int
	}{
		{3, false, []int{1, 2, 3, 4, 5}, []int{4, 5, 1, 2, 3}},
		{1, false, []int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}},
		{2, false, []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
		{5, false, []int{1, 2, 3}, []int{3, 1, 2}},
		{1, false, []int{}, []int{}},
	}

	for _, test := range tests {
		got := Rotate(test.by, test.data, test.right)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%v != %v", got, test.want)
		}
	}
}
