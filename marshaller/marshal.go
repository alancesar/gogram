package marshaller

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
	"reflect"
)

func Marshal(input measure.Measurable) ([]byte, error) {
	if isNumeric(input) {
		return []byte(input.String()), nil
	}

	formatted := fmt.Sprintf(`"%s"`, input.String())
	return []byte(formatted), nil
}

func isNumeric(input measure.Measurable) bool {
	switch reflect.ValueOf(input).Kind() {
	case reflect.Float32, reflect.Float64, reflect.Int, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}
