package unmarshaller

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
)

const quotes = '"'

func Unmarshal[T measure.Measurable](self *T, fromString func(input string) T, bytes []byte) error {
	raw, err := unquoteIfQuoted(string(bytes))
	if err != nil {
		return err
	}

	*self = fromString(raw)
	return nil
}

// Thanks https://github.com/shopspring/decimal
func unquoteIfQuoted(value interface{}) (string, error) {
	var bytes []byte

	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", value, value)
	}

	if len(bytes) > 2 && bytes[0] == quotes && bytes[len(bytes)-1] == quotes {
		bytes = bytes[1 : len(bytes)-1]
	}

	return string(bytes), nil
}
