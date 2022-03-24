package volume

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
)

const (
	MilliliterUnit measure.Unit = "ml"
	LiterUnit      measure.Unit = "l"
	GallonUnit     measure.Unit = "gal"

	millilitersInLiters = 1000
	gallonsInLiters     = 0.26417205235815
)

var (
	parsers = measure.ParseMap[Volume]{
		MilliliterUnit: NewFromMilliliter,
		LiterUnit:      NewFromLiter,
		GallonUnit:     NewFromGallon,
	}
)

type (
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
	if v.system == measure.MetricSystem {
		return v.metricString()
	}

	return v.imperialString()
}

func (v Volume) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, v.String())
	return []byte(formatted), nil
}

func (v *Volume) UnmarshalJSON(bytes []byte) error {
	*v = NewFromString(string(bytes))
	return nil
}

func (v Volume) metricString() string {
	if v.liters < 1 {
		return fmt.Sprintf("%.0f %s", v.Milliliters(), MilliliterUnit)
	}

	return fmt.Sprintf("%.2f %s", v.Liters(), LiterUnit)
}

func (v Volume) imperialString() string {
	return fmt.Sprintf("%.2f %s", v.Gallons(), GallonUnit)
}

func createFromMetric(liters float64) Volume {
	return Volume{
		system:  measure.MetricSystem,
		liters:  liters,
		gallons: liters * gallonsInLiters,
	}
}

func createFromImperial(gallons float64) Volume {
	return Volume{
		system:  measure.ImperialSystem,
		liters:  gallons / gallonsInLiters,
		gallons: gallons,
	}
}
