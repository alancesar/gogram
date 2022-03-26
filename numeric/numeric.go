package numeric

import "strconv"

const (
	bitSize   = 64
	formatter = 'f'
)

func Format(value float64, precision int) string {
	return strconv.FormatFloat(value, formatter, precision, bitSize)
}
