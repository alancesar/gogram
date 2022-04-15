package measury

import (
	"fmt"
	"github.com/alancesar/gogram/numeric"
)

type (
	Unit interface {
		~int64
		Symbol() string
	}

	Measurable[U Unit] interface {
		fmt.Stringer
		IsZero() bool
		In(unit U) float64
	}
)

func Stringer[U Unit](measure Measurable[U], unit U) string {
	if measure.IsZero() {
		return fmt.Sprintf("0%s", unit.Symbol())
	}

	value := numeric.Format(measure.In(unit))
	return fmt.Sprintf("%s%s", value, unit.Symbol())
}

func Marshaller[U Unit](input Measurable[U]) ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, input.String())
	return []byte(formatted), nil
}
