package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZero  Kelvin  = 0
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gºK", k) }

func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15) *9/5 + 32 ) }
func KToC(k Kelvin) Celsius { return Celsius(k + 273.15) }
