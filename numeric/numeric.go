package numeric

import "strconv"

const (
	formatter        = 'f'
	defaultPrecision = -1
	bitSize          = 64
)

func Format(value float64) string {
	return FormatWithPrecision(value, defaultPrecision)
}

func FormatWithPrecision(value float64, precision int) string {
	return strconv.FormatFloat(value, formatter, precision, bitSize)
}
