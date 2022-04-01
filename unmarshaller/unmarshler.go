package unmarshaller

import (
	"github.com/alancesar/gogram/measure"
	"strconv"
)

func Unmarshal[T measure.Measurable](self *T, stringer func(input string) T, bytes []byte) error {
	raw, err := strconv.Unquote(string(bytes))
	if err != nil {
		return err
	}

	*self = stringer(raw)
	return nil
}
