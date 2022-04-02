package marshaller

import (
	"fmt"
	"github.com/alancesar/gogram/measure"
	"reflect"
	"testing"
)

type (
	fakeNumberMeasure int
	fakeStringMeasure string
)

func (f fakeStringMeasure) String() string {
	return string(f)
}

func (f fakeStringMeasure) IsZero() bool {
	return len(f) == 0
}

func (f fakeNumberMeasure) String() string {
	return fmt.Sprintf("%d", f)
}

func (f fakeNumberMeasure) IsZero() bool {
	return f == 0
}

func TestMarshal(t *testing.T) {
	type args struct {
		input measure.Measurable
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should marshal as string",
			args: args{
				input: fakeStringMeasure("abc"),
			},
			want:    []byte(`"abc"`),
			wantErr: false,
		},
		{
			name: "Should marshal as number",
			args: args{
				input: fakeNumberMeasure(123),
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

func TestMarshalWithQuotes(t *testing.T) {
	type args struct {
		input measure.Measurable
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should marshal as string",
			args: args{
				input: fakeStringMeasure("abc"),
			},
			want:    []byte(`"abc"`),
			wantErr: false,
		},
		{
			name: "Should marshal as string event it's a number",
			args: args{
				input: fakeNumberMeasure(123),
			},
			want:    []byte(`"123"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalWithQuotes(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalWithQuotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalWithQuotes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNumeric(t *testing.T) {
	type args struct {
		input measure.Measurable
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true if it's a number",
			args: args{
				fakeNumberMeasure(123),
			},
			want: true,
		},
		{
			name: "Should return false if it isn't a number",
			args: args{
				fakeStringMeasure("123"),
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
