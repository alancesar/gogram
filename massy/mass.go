package massy

import (
	"github.com/alancesar/gogram/measury"
	"github.com/alancesar/gogram/numeric"
)

const (
	_             = iota
	Metric System = iota
	Imperial

	Milligram Unit = 1
	Gram           = Milligram * 1000
	Kilogram       = Gram * 1000
	Ounce          = Milligram * 28350
	Pound          = Ounce * 16

	defaultUnit   = Gram
	decimalPlaces = 2
)

type (
	Unit   int64
	System byte

	Mass struct {
		source     Unit
		milligrams int64
	}
)

func New(value float64, unit Unit) Mass {
	return Mass{
		source:     unit,
		milligrams: int64(value * float64(unit)),
	}
}

func (m Mass) IsZero() bool {
	return m.milligrams == 0
}

func (m Mass) In(unit Unit) float64 {
	value := float64(m.milligrams) / float64(unit)
	return numeric.Round(value, decimalPlaces)
}

func (m Mass) String() string {
	return measury.Stringer[Unit](m, m.findBestUnit())
}

func (m Mass) MarshalJSON() ([]byte, error) {
	return measury.Marshaller[Unit](m)
}

func (m Mass) findBestUnit() Unit {
	if m.source.System() == Metric {
		return m.findBestUnitForMetric()
	}

	return m.findBestUnitForImperial()
}

func (m Mass) findBestUnitForMetric() Unit {
	switch {
	case m.milligrams >= int64(Kilogram):
		return Kilogram
	case m.milligrams >= int64(Gram):
		return Gram
	default:
		return Milligram
	}
}

func (m Mass) findBestUnitForImperial() Unit {
	switch {
	case m.milligrams >= int64(Pound):
		return Pound
	default:
		return Ounce
	}
}

func (u Unit) System() System {
	switch u {
	case Milligram, Kilogram, Gram:
		return Metric
	case Ounce, Pound:
		return Imperial
	default:
		return defaultUnit.System()
	}
}

func (u Unit) Symbol() string {
	switch u {
	case Milligram:
		return "mg"
	case Gram:
		return "g"
	case Kilogram:
		return "kg"
	case Ounce:
		return "oz"
	case Pound:
		return "lb"
	default:
		return defaultUnit.Symbol()
	}
}
