package temperature

import (
	"github.com/alancesar/gogram/measure"
	"reflect"
	"testing"
)

func TestNewFromCelsius(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Temperature
	}{
		{
			name: "Should parse from Celsius",
			args: args{
				value: 50,
			},
			want: Temperature{
				unit:       CelsiusUnit,
				celsius:    50,
				fahrenheit: 122,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromCelsius(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromCelsius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromFahrenheit(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Temperature
	}{
		{
			name: "Should parse from Fahrenheit",
			args: args{
				value: 68,
			},
			want: Temperature{
				unit:       FahrenheitUnit,
				celsius:    20,
				fahrenheit: 68,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromFahrenheit(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromFahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want Temperature
	}{
		{
			name: "Should parse from '1C' string",
			args: args{
				input: "1C",
			},
			want: NewFromCelsius(1),
		},
		{
			name: "Should parse from '1ºC' string",
			args: args{
				input: "1ºC",
			},
			want: NewFromCelsius(1),
		},
		{
			name: "Should parse from '1°C' string",
			args: args{
				input: "1°C",
			},
			want: NewFromCelsius(1),
		},
		{
			name: "Should parse from '1F' string",
			args: args{
				input: "1F",
			},
			want: NewFromFahrenheit(1),
		},
		{
			name: "Should parse from '1ºF' string",
			args: args{
				input: "1ºF",
			},
			want: NewFromFahrenheit(1),
		},
		{
			name: "Should parse from '1°F' string",
			args: args{
				input: "1°F",
			},
			want: NewFromFahrenheit(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromString(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_IsZero(t1 *testing.T) {
	type fields struct {
		celsius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Should return true if is empty",
			fields: fields{
				celsius: 1,
			},
			want: true,
		},
		{
			name: "Should return false if is not empty",
			fields: fields{
				celsius: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Temperature{
				celsius: tt.fields.celsius,
			}
			if got := t.IsZero(); got != tt.want {
				t1.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_Celsius(t1 *testing.T) {
	type fields struct {
		unit    measure.Unit
		celsius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should get Celsius value properly",
			fields: fields{
				celsius: 15,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Temperature{
				unit:    tt.fields.unit,
				celsius: tt.fields.celsius,
			}
			if got := t.Celsius(); got != tt.want {
				t1.Errorf("Celsius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_Fahrenheit(t1 *testing.T) {
	type fields struct {
		fahrenheit float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should get Fahrenheit value properly",
			fields: fields{
				fahrenheit: 15,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Temperature{
				fahrenheit: tt.fields.fahrenheit,
			}
			if got := t.Fahrenheit(); got != tt.want {
				t1.Errorf("Fahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_String(t1 *testing.T) {
	type fields struct {
		unit       measure.Unit
		celsius    float64
		fahrenheit float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should print 15°C",
			fields: fields{
				unit:    CelsiusUnit,
				celsius: 15,
			},
			want: "15°C",
		},
		{
			name: "Should print 15°F",
			fields: fields{
				unit:       FahrenheitUnit,
				fahrenheit: 15,
			},
			want: "15°F",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Temperature{
				unit:       tt.fields.unit,
				celsius:    tt.fields.celsius,
				fahrenheit: tt.fields.fahrenheit,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_MarshalJSON(t1 *testing.T) {
	type fields struct {
		celsius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Should marshal properly",
			fields: fields{
				celsius: 15,
			},
			want:    []byte(`15°C`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewFromCelsius(tt.fields.celsius)
			got, err := t.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t1.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Temperature
		wantErr bool
	}{
		{
			name:    "Should unmarshal properly",
			args:    args{},
			want:    NewFromCelsius(23),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			temp := &Temperature{}
			if err := temp.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t1.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(*temp, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", temp, tt.want)
			}
		})
	}
}
