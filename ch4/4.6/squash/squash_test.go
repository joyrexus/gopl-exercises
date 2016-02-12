package squash_test

import (
	"bytes"
	"testing"

	. "github.com/joyrexus/gopl-exercises/ch4/4.6/squash"
)

type Cases []struct {
	input []byte
	want  []byte
}

var cases = Cases{
	{[]byte(""), []byte("")},
	{[]byte("a"), []byte("a")},
	{[]byte("foo bar"), []byte("foo bar")},
	{[]byte("foo\tbar"), []byte("foo bar")},
	{[]byte("\tfoo\t\vbar\r\n"), []byte(" foo bar ")},
	{[]byte("x  z"), []byte("x z")},
	{[]byte("foo\t\tbar"), []byte("foo bar")},
	{[]byte("世  z"), []byte("世 z")},
	{[]byte("x\t\t\t世\t\t\ty \v\v\n  z"), []byte("x 世 y z")},
}

func TestSquash(t *testing.T) {
	var tests Cases
	copy(tests, cases) // use a copy since each case gets modified in-place
	for _, c := range tests {
		if got := Squash(c.input); !bytes.Equal(got, c.want) {
			t.Errorf("got %s, want %s", got, c.want)
		}
	}
}

func TestTrim(t *testing.T) {
	var tests Cases
	copy(tests, cases) // use a copy since each case gets modified in-place
	for _, c := range tests {
		if got := Trim(c.input); !bytes.Equal(got, c.want) {
			t.Errorf("got %s, want %s", got, c.want)
		}
	}
}
