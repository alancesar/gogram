package marshaller

import (
	"encoding/json"
	"fmt"
	"github.com/alancesar/gogram/measure"
	"reflect"
)

func Marshal(input measure.Measurable) ([]byte, error) {
	if isNumeric(input) {
		return json.Marshal(input)
	}

	if stringer, ok := input.(fmt.Stringer); ok {
		formatted := fmt.Sprintf(`"%s"`, stringer.String())
		return []byte(formatted), nil
	}

	return json.Marshal(input)
}

func isNumeric(input measure.Measurable) bool {
	switch reflect.ValueOf(input).Kind() {
	case reflect.Float32, reflect.Float64, reflect.Int, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}
