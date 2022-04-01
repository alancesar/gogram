package marshaller

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
)

func Marshal(input measure.Measurable) ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, input.String())
	return []byte(formatted), nil
}
