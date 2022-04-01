package measure

//go:generate stringer -type=System

import (
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
	Parser[T Measurable]    func(value float64) T
	ParserMap[T Measurable] map[string]Parser[T]

	System int

	Measurable interface {
		fmt.Stringer
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
