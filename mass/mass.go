package mass

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
	"github.com/alancesar/gogram/numeric"
)

const (
	Milligram Unit = "mg"
	Gram      Unit = "g"
	Kilogram  Unit = "kg"
	Pound     Unit = "lb"
	Ounce     Unit = "oz"

	milligramsInGrams = 1000
	gramsInKilograms  = 1000
	kilogramsInGrams  = 0.001
	poundsInGrams     = 453.592
	poundsInOunces    = 16
)

var (
	parsers = measure.ParseMap[Mass]{
		measure.Unit(Milligram): NewFromMilligram,
		measure.Unit(Gram):      NewFromGram,
		measure.Unit(Kilogram):  NewFromKilogram,
		measure.Unit(Pound):     NewFromPound,
		measure.Unit(Ounce):     NewFromOunce,
	}
)

type (
	Unit measure.Unit

	Mass struct {
		system        measure.System
		grams, pounds float64
	}
)

func NewFromString(input string) Mass {
	return parsers.Parse(input)
}

func NewFromMilligram(value float64) Mass {
	return createFromMetric(value / milligramsInGrams)
}

func NewFromGram(value float64) Mass {
	return createFromMetric(value)
}

func NewFromKilogram(value float64) Mass {
	return createFromMetric(value / kilogramsInGrams)
}

func NewFromPound(value float64) Mass {
	return createFromImperial(value)
}

func NewFromOunce(value float64) Mass {
	return createFromImperial(value / poundsInOunces)
}

func (m Mass) IsZero() bool {
	return m.grams == 0 && m.pounds == 0
}

func (m Mass) Milligrams() float64 {
	return m.grams * milligramsInGrams
}

func (m Mass) Grams() float64 {
	return m.grams
}

func (m Mass) Kilograms() float64 {
	return m.grams * kilogramsInGrams
}

func (m Mass) Pounds() float64 {
	return m.pounds
}

func (m Mass) Ounces() float64 {
	return m.pounds * poundsInOunces
}

func (m Mass) String() string {
	unit := m.findBestUnit()
	return m.StringIn(unit)
}

func (m Mass) StringIn(unit Unit) string {
	value, err := m.Float64In(unit)
	if err != nil {
		return ""
	}
	formatted := numeric.Format(value)
	return fmt.Sprintf("%s %s", formatted, unit)
}

func (m Mass) Float64In(unit Unit) (float64, error) {
	switch unit {
	case Milligram:
		return m.Milligrams(), nil
	case Gram:
		return m.Grams(), nil
	case Kilogram:
		return m.Kilograms(), nil
	case Pound:
		return m.Pounds(), nil
	case Ounce:
		return m.Ounces(), nil
	default:
		return 0, fmt.Errorf("%s is an invalid unit for mass", unit)
	}
}

func (m Mass) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, m.String())
	return []byte(formatted), nil
}

func (m *Mass) UnmarshalJSON(bytes []byte) error {
	*m = NewFromString(string(bytes))
	return nil
}

func (m Mass) findBestUnit() Unit {
	if m.system == measure.Metric {
		switch {
		case m.grams >= gramsInKilograms:
			return Kilogram
		case m.grams < 1:
			return Milligram
		default:
			return Gram
		}
	}

	switch {
	case m.pounds < 1:
		return Ounce
	default:
		return Pound
	}
}

func createFromMetric(grams float64) Mass {
	return Mass{
		system: measure.Metric,
		grams:  grams,
		pounds: grams / poundsInGrams,
	}
}

func createFromImperial(pounds float64) Mass {
	return Mass{
		system: measure.Imperial,
		grams:  pounds * poundsInGrams,
		pounds: pounds,
	}
}
