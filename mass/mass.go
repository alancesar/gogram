package mass

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
)

const (
	MilligramUnit measure.Unit = "mg"
	GramUnit      measure.Unit = "g"
	KilogramUnit  measure.Unit = "kg"
	PoundUnit     measure.Unit = "lb"
	OunceUnit     measure.Unit = "oz"

	milligramsInGrams = 1000
	gramsInKilograms  = 1000
	kilogramsInGrams  = 0.001
	poundsInGrams     = 453.592
	poundsInOunces    = 16
)

var (
	parsers = measure.ParseMap[Mass]{
		MilligramUnit: NewFromMilligram,
		GramUnit:      NewFromGram,
		KilogramUnit:  NewFromKilogram,
		PoundUnit:     NewFromPound,
		OunceUnit:     NewFromOunce,
	}
)

type (
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

func (m Mass) String() (formatted string) {
	if m.system == measure.ImperialSystem {
		return m.imperialString()
	}

	return m.metricString()
}

func (m Mass) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, m.String())
	return []byte(formatted), nil
}

func (m *Mass) UnmarshalJSON(bytes []byte) error {
	*m = NewFromString(string(bytes))
	return nil
}

func (m Mass) metricString() string {
	switch {
	case m.grams >= gramsInKilograms:
		return fmt.Sprintf("%.2f %s", m.Kilograms(), KilogramUnit)
	case m.grams < 1:
		return fmt.Sprintf("%.0f %s", m.Milligrams(), MilligramUnit)
	default:
		return fmt.Sprintf("%.2f %s", m.Grams(), GramUnit)
	}
}

func (m Mass) imperialString() string {
	if m.pounds < 1 {
		return fmt.Sprintf("%.f %s", m.Ounces(), OunceUnit)
	}

	return fmt.Sprintf("%.2f %s", m.Pounds(), PoundUnit)
}

func createFromMetric(grams float64) Mass {
	return Mass{
		system: measure.MetricSystem,
		grams:  grams,
		pounds: grams / poundsInGrams,
	}
}

func createFromImperial(pounds float64) Mass {
	return Mass{
		system: measure.ImperialSystem,
		grams:  pounds * poundsInGrams,
		pounds: pounds,
	}
}
