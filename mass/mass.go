package mass

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
	"strconv"
)

const (
	MilligramUnit Unit = "mg"
	GramUnit      Unit = "g"
	KilogramUnit  Unit = "kg"
	PoundUnit     Unit = "lb"
	OunceUnit     Unit = "oz"

	milligramsInGrams = 1000
	gramsInKilograms  = 1000
	kilogramsInGrams  = 0.001
	poundsInGrams     = 453.592
	poundsInOunces    = 16

	bitSize   = 64
	formatter = 'f'
)

var (
	MilligramPrecision = 0
	GramPrecision      = 2
	KilogramPrecision  = 2
	PoundPrecision     = 2
	OuncePrecision     = 0

	precisions = map[Unit]*int{
		MilligramUnit: &MilligramPrecision,
		GramUnit:      &GramPrecision,
		KilogramUnit:  &KilogramPrecision,
		PoundUnit:     &PoundPrecision,
		OunceUnit:     &OuncePrecision,
	}

	parsers = measure.ParseMap[Mass]{
		measure.Unit(MilligramUnit): NewFromMilligram,
		measure.Unit(GramUnit):      NewFromGram,
		measure.Unit(KilogramUnit):  NewFromKilogram,
		measure.Unit(PoundUnit):     NewFromPound,
		measure.Unit(OunceUnit):     NewFromOunce,
	}

	getters = map[Unit]func(mass Mass) float64{
		MilligramUnit: func(mass Mass) float64 {
			return mass.Milligrams()
		},
		GramUnit: func(mass Mass) float64 {
			return mass.Grams()
		},
		KilogramUnit: func(mass Mass) float64 {
			return mass.Kilograms()
		},
		PoundUnit: func(mass Mass) float64 {
			return mass.Pounds()
		},
		OunceUnit: func(mass Mass) float64 {
			return mass.Ounces()
		},
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
	getter, ok := getters[unit]
	if !ok {
		return ""
	}

	formatted := format(getter(m), *precisions[unit])
	return fmt.Sprintf("%s %s", formatted, unit)
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
			return KilogramUnit
		case m.grams < 1:
			return MilligramUnit
		default:
			return GramUnit
		}
	}

	switch {
	case m.pounds < 1:
		return OunceUnit
	default:
		return PoundUnit
	}
}

func format(value float64, precision int) string {
	return strconv.FormatFloat(value, formatter, precision, bitSize)
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
