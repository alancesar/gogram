package measure

//go:generate stringer -type=System

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	Metric System = iota
	Imperial

	valueIndex = 1
	unitIndex  = 3
)

var (
	regex = regexp.MustCompile(`(\d*\.?\d*)(\s?)(\D{1,7})`)
)

type (
	Parser[T any]   func(value float64) T
	ParseMap[T any] map[Unit]Parser[T]
	Unit            string
	System          int

	Measurable interface {
		fmt.Stringer
		json.Marshaler
		json.Unmarshaler
	}
)

func (m ParseMap[T]) Parse(input string) T {
	var empty T

	elements := regex.FindStringSubmatch(input)
	if elements == nil {
		return empty
	}

	unit := strings.ToLower(elements[unitIndex])
	value, _ := strconv.ParseFloat(elements[valueIndex], 64)

	if builder, ok := m[Unit(unit)]; !ok {
		return empty
	} else {
		return builder(value)
	}
}
