package measure

import (
	"fmt"
	"reflect"
	"testing"
)

type fakeMeasurable string

func (f fakeMeasurable) IsZero() bool {
	return f != ""
}

var (
	parseFn = func(value float64) fakeMeasurable {
		return fakeMeasurable(fmt.Sprintf("%.2f", value))
	}
)

func TestBuilderMap_Parse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		m    ParserMap[fakeMeasurable]
		args args
		want fakeMeasurable
	}{
		{
			name: "Should parse properly",
			m: ParserMap[fakeMeasurable]{
				"foo": parseFn,
			},
			args: args{
				input: "16 foo",
			},
			want: "16.00",
		},
		{
			name: "Should return empty if have an invalid value",
			m: ParserMap[fakeMeasurable]{
				"foo": parseFn,
			},
			args: args{
				input: "x foo",
			},
			want: "",
		},
		{
			name: "Should return empty if is an invalid pattern",
			m: ParserMap[fakeMeasurable]{
				"foo": parseFn,
			},
			args: args{
				input: "bar 16",
			},
			want: "",
		},
		{
			name: "Should return empty if have no symbol",
			m: ParserMap[fakeMeasurable]{
				"foo": parseFn,
			},
			args: args{
				input: "16",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Parse(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
