package measure

//go:generate stringer -type=System

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	Metric System = iota
	Imperial

	valueIndex = 1
	unitIndex  = 3

	quotes = '"'
)

var (
	regex = regexp.MustCompile(`(\d*\.?\d*)(\s?)(\D{1,7})`)
)

type (
	Parser[T Measurable]    func(value float64) T
	ParserMap[T Measurable] map[string]Parser[T]

	System int

	Measurable interface {
		IsZero() bool
	}
)

func (m ParserMap[T]) Parse(input string) T {
	var empty T

	elements := regex.FindStringSubmatch(input)
	if elements == nil {
		return empty
	}

	unit := strings.ToLower(elements[unitIndex])
	value, _ := strconv.ParseFloat(elements[valueIndex], 64)

	if builder, ok := m[unit]; !ok {
		return empty
	} else {
		return builder(value)
	}
}

func Marshal(input Measurable) ([]byte, error) {
	if isNumeric(input) {
		return json.Marshal(input)
	}

	if stringer, ok := input.(fmt.Stringer); ok {
		formatted := fmt.Sprintf(`"%s"`, stringer.String())
		return []byte(formatted), nil
	}

	return json.Marshal(input)
}

func Unmarshal[T Measurable](self *T, fromString func(input string) T, bytes []byte) error {
	raw, err := unquoteIfQuoted(string(bytes))
	if err != nil {
		return err
	}

	*self = fromString(raw)
	return nil
}

func isNumeric(input Measurable) bool {
	switch reflect.ValueOf(input).Kind() {
	case reflect.Float32, reflect.Float64, reflect.Int, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
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
