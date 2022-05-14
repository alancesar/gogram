package json

import (
	"fmt"
	"strconv"
)

func Marshal(input fmt.Stringer) ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, input.String())
	return []byte(formatted), nil
}

func UnmarshalQuoted[T any](self *T, fromString func(input string) T, bytes []byte) error {
	raw, err := strconv.Unquote(string(bytes))
	if err != nil {
		return err
	}

	return Unmarshal(self, fromString, []byte(raw))
}

func Unmarshal[T any](self *T, fromString func(input string) T, bytes []byte) error {
	*self = fromString(string(bytes))
	return nil
}
