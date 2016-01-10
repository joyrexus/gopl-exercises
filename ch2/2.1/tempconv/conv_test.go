package tempconv

import (
	"fmt"
	"strconv"
	"testing"
)

func round(x fmt.Stringer) (float64, error) {
	r := fmt.Sprintf("%0.3f", x)
	f, err := strconv.ParseFloat(r, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// Ensure we can convert Fahrenheit to Kelvin.
func TestFToK(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{100, 310.928},
		{0, 255.372},
		{-459.67, 0},
		{-279.67, 100},
	}

	for _, c := range cases {
		f := Fahrenheit(c.in)
		got, err := round(FToK(f))
		if err != nil {
			panic(err)
		}
		if got != c.want {
			t.Errorf("FToK(%f) == %f, want %f", c.in, got, c.want)
		}
	}
}

// Ensure we can convert Kelvin to Fahrenheit.
func TestKToF(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{310.928, 100},
		{255.372, 0},
		{0, -459.67},
		{100, -279.67},
	}

	for _, c := range cases {
		k := Kelvin(c.in)
		got, err := round(KToF(k))
		if err != nil {
			panic(err)
		}
		if got != c.want {
			t.Errorf("FToK(%f) == %f, want %f", c.in, got, c.want)
		}
	}
}
