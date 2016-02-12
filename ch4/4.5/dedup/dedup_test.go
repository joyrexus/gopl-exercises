package dedup_test

import (
	"reflect"
	"testing"

	. "github.com/joyrexus/gopl-exercises/ch4/4.5/dedup"
)

func TestDedup(t *testing.T) {

	tests := []struct {
		data []string
		want []string
	}{
		{
			[]string{}, 
			[]string{},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			[]string{"a", "a", "b", "a", "c", "b", "b", "b"},
			[]string{"a", "b", "a", "c", "b"},
		},
		{
			[]string{"a", "b", "b", "a", "a", "c", "c", "b", "b"},
			[]string{"a", "b", "a", "c", "b"},
		},
	}

	for _, test := range tests {
		got := Dedup(test.data)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%v != %v", got, test.want)
		}
	}
}
