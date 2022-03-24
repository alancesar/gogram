package temperature

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
)

const (
	CelsiusUnit    measure.Unit = "°C"
	FahrenheitUnit measure.Unit = "°F"
)

var (
	parsers = measure.ParseMap[Temperature]{
		"c":  NewFromCelsius,
		"ºc": NewFromCelsius,
		"°c": NewFromCelsius,
		"f":  NewFromFahrenheit,
		"ºf": NewFromFahrenheit,
		"°f": NewFromFahrenheit,
	}
)

type (
	Temperature struct {
		unit                measure.Unit
		celsius, fahrenheit float64
	}
)

func NewFromString(input string) Temperature {
	return parsers.Parse(input)
}

func NewFromCelsius(value float64) Temperature {
	return Temperature{
		unit:       CelsiusUnit,
		celsius:    value,
		fahrenheit: (value * 1.8) + 32,
	}
}

func NewFromFahrenheit(value float64) Temperature {
	return Temperature{
		unit:       FahrenheitUnit,
		celsius:    (value - 32) / 1.8,
		fahrenheit: value,
	}
}

func (t Temperature) IsZero() bool {
	return t.celsius == 0 && t.fahrenheit == 0
}

func (t Temperature) Celsius() float64 {
	return t.celsius
}

func (t Temperature) Fahrenheit() float64 {
	return t.fahrenheit
}

func (t Temperature) String() string {
	if t.unit == CelsiusUnit {
		return fmt.Sprintf("%.f%s", t.Celsius(), CelsiusUnit)
	}

	return fmt.Sprintf("%.f%s", t.Fahrenheit(), FahrenheitUnit)
}

func (t Temperature) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, t.String())
	return []byte(formatted), nil
}

func (t *Temperature) UnmarshalJSON(bytes []byte) error {
	*t = NewFromString(string(bytes))
	return nil
}
