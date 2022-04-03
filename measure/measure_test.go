package measure

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	fakeMeasurable       string
	fakeStringMeasurable string
	fakeNumberMeasurable int
)

func (f fakeMeasurable) IsZero() bool {
	return f != ""
}

func (f fakeStringMeasurable) String() string {
	return fmt.Sprintf("%s implements Stringer", string(f))
}

func (f fakeStringMeasurable) IsZero() bool {
	return f != ""
}

func (f fakeNumberMeasurable) IsZero() bool {
	return f == 0
}

var (
	parseFn = func(value float64) fakeStringMeasurable {
		return fakeStringMeasurable(fmt.Sprintf("%.2f", value))
	}
)

func TestBuilderMap_Parse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		m    ParserMap[fakeStringMeasurable]
		args args
		want fakeStringMeasurable
	}{
		{
			name: "Should parse properly",
			m: ParserMap[fakeStringMeasurable]{
				"foo": parseFn,
			},
			args: args{
				input: "16 foo",
			},
			want: "16.00",
		},
		{
			name: "Should return empty if have an invalid value",
			m: ParserMap[fakeStringMeasurable]{
				"foo": parseFn,
			},
			args: args{
				input: "x foo",
			},
			want: "",
		},
		{
			name: "Should return empty if is an invalid pattern",
			m: ParserMap[fakeStringMeasurable]{
				"foo": parseFn,
			},
			args: args{
				input: "bar 16",
			},
			want: "",
		},
		{
			name: "Should return empty if have no symbol",
			m: ParserMap[fakeStringMeasurable]{
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

func TestMarshal(t *testing.T) {
	type args struct {
		input Measurable
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should marshal as string using String() method",
			args: args{
				input: fakeStringMeasurable("abc"),
			},
			want:    []byte(`"abc implements Stringer"`),
			wantErr: false,
		},
		{
			name: "Should marshal as string",
			args: args{
				input: fakeMeasurable("abc"),
			},
			want:    []byte(`"abc"`),
			wantErr: false,
		},
		{
			name: "Should marshal as number",
			args: args{
				input: fakeNumberMeasurable(123),
			},
			want:    []byte("123"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNumeric(t *testing.T) {
	type args struct {
		input Measurable
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true if it's a number",
			args: args{
				fakeNumberMeasurable(123),
			},
			want: true,
		},
		{
			name: "Should return false if it isn't a number",
			args: args{
				fakeStringMeasurable("123"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumeric(tt.args.input); got != tt.want {
				t.Errorf("isNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
