package reverse_test

import (
	"bytes"
	"testing"

	. "github.com/joyrexus/gopl-exercises/ch4/4.7/reverse"
)

var tests = []struct {
	input []byte
	want  []byte
}{
	{[]byte(""), []byte("")},
	{[]byte("a"), []byte("a")},
	{[]byte("foo bar"), []byte("rab oof")},
	{[]byte("Hello, 世界"), []byte("界世 ,olleH")},
}

func TestReverse(t *testing.T) {
	for _, test := range tests {
		if Reverse(test.input); !bytes.Equal(test.input, test.want) {
			t.Errorf("got %s, want %s", test.input, test.want)
		}
	}
}
