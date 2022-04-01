package temperature

import (
	"fmt"
	"github.com/alancesar/gogram/marshaller"
	"github.com/alancesar/gogram/measure"
	"github.com/alancesar/gogram/numeric"
	"github.com/alancesar/gogram/unmarshaller"
)

const (
	Celsius    Unit = "°C"
	Fahrenheit Unit = "°F"
)

var (
	parsers = measure.ParserMap[Temperature]{
		"c":  NewFromCelsius,
		"ºc": NewFromCelsius,
		"°c": NewFromCelsius,
		"f":  NewFromFahrenheit,
		"ºf": NewFromFahrenheit,
		"°f": NewFromFahrenheit,
	}
)

type (
	Unit string

	Temperature struct {
		unit                Unit
		celsius, fahrenheit float64
	}
)

func NewFromString(input string) Temperature {
	return parsers.Parse(input)
}

func NewFromCelsius(value float64) Temperature {
	return Temperature{
		unit:       Celsius,
		celsius:    value,
		fahrenheit: (value * 1.8) + 32,
	}
}

func NewFromFahrenheit(value float64) Temperature {
	return Temperature{
		unit:       Fahrenheit,
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
	unit := t.findBestUnit()
	return t.StringIn(unit)
}

func (t Temperature) StringIn(unit Unit) string {
	value, err := t.Float64In(unit)
	if err != nil {
		return ""
	}
	formatted := numeric.Format(value)
	return fmt.Sprintf("%s%s", formatted, unit)
}

func (t Temperature) Float64In(unit Unit) (float64, error) {
	switch unit {
	case Celsius:
		return t.Celsius(), nil
	case Fahrenheit:
		return t.Fahrenheit(), nil
	default:
		return 0, fmt.Errorf("%s is an invalid unit for temperature", unit)
	}
}

func (t Temperature) MarshalJSON() ([]byte, error) {
	return marshaller.Marshal(t)
}

func (t *Temperature) UnmarshalJSON(bytes []byte) error {
	return unmarshaller.Unmarshal(t, NewFromString, bytes)
}

func (t Temperature) findBestUnit() Unit {
	if t.unit == Celsius {
		return Celsius
	}

	return Fahrenheit
}
