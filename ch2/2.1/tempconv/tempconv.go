// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Kelvin float64
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
