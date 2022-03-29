package volume

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
	"github.com/alancesar/gogram/numeric"
	"strconv"
)

const (
	Milliliter Unit = "ml"
	Liter      Unit = "l"
	Gallon     Unit = "gal"
	Ounce      Unit = "fl. Oz"

	millilitersInLiters = 1000
	litersInGallons     = 4.54609
	ouncesInGallons     = 160
)

var (
	parsers = measure.ParserMap[Volume]{
		"ml":     NewFromMilliliter,
		"l":      NewFromLiter,
		"gal":    NewFromGallon,
		"fl. oz": NewFromOunce,
		"fl oz":  NewFromOunce,
	}
)

type (
	Unit string

	Volume struct {
		system          measure.System
		liters, gallons float64
	}
)

func NewFromString(input string) Volume {
	return parsers.Parse(input)
}

func NewFromMilliliter(value float64) Volume {
	return createFromMetric(value / millilitersInLiters)
}

func NewFromLiter(value float64) Volume {
	return createFromMetric(value)
}

func NewFromGallon(value float64) Volume {
	return createFromImperial(value)
}

func NewFromOunce(value float64) Volume {
	return createFromImperial(value / ouncesInGallons)
}

func (v Volume) IsZero() bool {
	return v.liters == 0 && v.gallons == 0
}

func (v Volume) Milliliters() float64 {
	return v.liters * millilitersInLiters
}

func (v Volume) Liters() float64 {
	return v.liters
}

func (v Volume) Gallons() float64 {
	return v.gallons
}

func (v Volume) Ounces() float64 {
	return v.gallons * ouncesInGallons
}

func (v Volume) String() string {
	unit := v.findBestUnit()
	return v.StringIn(unit)
}

func (v Volume) StringIn(unit Unit) string {
	value, err := v.Float64In(unit)
	if err != nil {
		return ""
	}
	formatted := numeric.Format(value)
	return fmt.Sprintf("%s %s", formatted, unit)
}

func (v Volume) Float64In(unit Unit) (float64, error) {
	switch unit {
	case Milliliter:
		return v.Milliliters(), nil
	case Liter:
		return v.Liters(), nil
	case Gallon:
		return v.Gallons(), nil
	case Ounce:
		return v.Ounces(), nil
	default:
		return 0, fmt.Errorf("%s is an invalid unit for volume", unit)
	}
}

func (v Volume) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, v.String())
	return []byte(formatted), nil
}

func (v *Volume) UnmarshalJSON(bytes []byte) error {
	raw, err := strconv.Unquote(string(bytes))
	if err != nil {
		return err
	}

	*v = NewFromString(raw)
	return nil
}

func (v Volume) findBestUnit() Unit {
	if v.system == measure.Metric {
		switch {
		case v.liters < 1:
			return Milliliter
		default:
			return Liter
		}
	}

	return Gallon
}

func createFromMetric(liters float64) Volume {
	return Volume{
		system:  measure.Metric,
		liters:  liters,
		gallons: liters / litersInGallons,
	}
}

func createFromImperial(gallons float64) Volume {
	return Volume{
		system:  measure.Imperial,
		liters:  gallons * litersInGallons,
		gallons: gallons,
	}
}
