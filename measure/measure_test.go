package measure

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuilderMap_Parse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		m    ParseMap[string]
		args args
		want string
	}{
		{
			name: "Should parse properly",
			m: ParseMap[string]{
				"foo": func(value float64) string {
					return fmt.Sprintf("%.2f", value)
				},
			},
			args: args{
				input: "16 foo",
			},
			want: "16.00",
		},
		{
			name: "Should return empty if have an invalid value",
			m: ParseMap[string]{
				"foo": func(value float64) string {
					return fmt.Sprintf("%.2f", value)
				},
			},
			args: args{
				input: "x foo",
			},
			want: "",
		},
		{
			name: "Should return empty if is an invalid pattern",
			m: ParseMap[string]{
				"foo": func(value float64) string {
					return fmt.Sprintf("%.2f", value)
				},
			},
			args: args{
				input: "bar 16",
			},
			want: "",
		},
		{
			name: "Should return empty if have no symbol",
			m: ParseMap[string]{
				"foo": func(value float64) string {
					return fmt.Sprintf("%.2f", value)
				},
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
