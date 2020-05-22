package tempconv

import "fmt"

type Kelvin float64
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

func (k Kelvin) String() string           { return fmt.Sprintf("%g K", k) }
func (k Kelvin) ToCelsius() Celsius       { return Celsius(k - 273.15) }
func (k Kelvin) ToFahrenheit() Fahrenheit { return k.ToCelsius().ToFahrenheit() }

func (c Celsius) String() string           { return fmt.Sprintf("%g℃", c) }
func (c Celsius) ToKelvin() Kelvin         { return Kelvin(c + 273.15) }
func (c Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func (f Fahrenheit) String() string     { return fmt.Sprintf("%g F", f) }
func (f Fahrenheit) ToKelvin() Kelvin   { return f.ToCelsius().ToKelvin() }
func (f Fahrenheit) ToCelsius() Celsius { return Celsius((f - 32) * 5 / 9) }
