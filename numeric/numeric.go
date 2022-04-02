package numeric

import (
	"math"
	"strconv"
)

const (
	formatter        = 'f'
	defaultPrecision = -1
	bitSize          = 64
)

func Format(precision float64) string {
	return FormatWithPrecision(precision, defaultPrecision)
}

func FormatWithPrecision(value float64, precision int) string {
	return strconv.FormatFloat(value, formatter, precision, bitSize)
}

func Round(input float64, decimal int) float64 {
	factor := math.Pow10(decimal)
	rounded := math.Round(input * factor)
	return float64(int(rounded)) / factor
}
