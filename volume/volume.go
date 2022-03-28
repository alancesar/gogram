package volume

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
	"github.com/alancesar/gogram/numeric"
)

const (
	MilliliterUnit Unit = "ml"
	LiterUnit      Unit = "l"
	GallonUnit     Unit = "gal"

	millilitersInLiters = 1000
	gallonsInLiters     = 0.26417205235815
)

var (
	parsers = measure.ParseMap[Volume]{
		measure.Unit(MilliliterUnit): NewFromMilliliter,
		measure.Unit(LiterUnit):      NewFromLiter,
		measure.Unit(GallonUnit):     NewFromGallon,
	}
)

type (
	Unit measure.Unit

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
	case MilliliterUnit:
		return v.Milliliters(), nil
	case LiterUnit:
		return v.Liters(), nil
	case GallonUnit:
		return v.Gallons(), nil
	default:
		return 0, fmt.Errorf("%s is an invalid unit for volume", unit)
	}
}

func (v Volume) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, v.String())
	return []byte(formatted), nil
}

func (v *Volume) UnmarshalJSON(bytes []byte) error {
	*v = NewFromString(string(bytes))
	return nil
}

func (v Volume) findBestUnit() Unit {
	if v.system == measure.Metric {
		switch {
		case v.liters < 1:
			return MilliliterUnit
		default:
			return LiterUnit
		}
	}

	return GallonUnit
}

func createFromMetric(liters float64) Volume {
	return Volume{
		system:  measure.Metric,
		liters:  liters,
		gallons: liters * gallonsInLiters,
	}
}

func createFromImperial(gallons float64) Volume {
	return Volume{
		system:  measure.Imperial,
		liters:  gallons / gallonsInLiters,
		gallons: gallons,
	}
}
