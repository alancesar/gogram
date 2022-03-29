package temperature

import (
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
				unit:       Celsius,
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
				unit:       Fahrenheit,
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

func TestTemperature_IsZero(t *testing.T) {
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
				celsius: 0,
			},
			want: true,
		},
		{
			name: "Should return false if is not empty",
			fields: fields{
				celsius: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temperature := Temperature{
				celsius: tt.fields.celsius,
			}
			if got := temperature.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_Celsius(t *testing.T) {
	type fields struct {
		unit    Unit
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
		t.Run(tt.name, func(t *testing.T) {
			temperature := Temperature{
				unit:    tt.fields.unit,
				celsius: tt.fields.celsius,
			}
			if got := temperature.Celsius(); got != tt.want {
				t.Errorf("Celsius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_Fahrenheit(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			temperature := Temperature{
				fahrenheit: tt.fields.fahrenheit,
			}
			if got := temperature.Fahrenheit(); got != tt.want {
				t.Errorf("Fahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_String(t *testing.T) {
	type fields struct {
		unit       Unit
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
				unit:    Celsius,
				celsius: 15,
			},
			want: "15°C",
		},
		{
			name: "Should print 15°F",
			fields: fields{
				unit:       Fahrenheit,
				fahrenheit: 15,
			},
			want: "15°F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temperature := Temperature{
				unit:       tt.fields.unit,
				celsius:    tt.fields.celsius,
				fahrenheit: tt.fields.fahrenheit,
			}
			if got := temperature.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_MarshalJSON(t *testing.T) {
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
			want:    []byte(`"15°C"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temperature := NewFromCelsius(tt.fields.celsius)
			got, err := temperature.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
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
			name: "Should unmarshal properly",
			args: args{
				bytes: []byte(`"23ºC"`),
			},
			want:    NewFromCelsius(23),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temp := &Temperature{}
			if err := temp.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(*temp, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", temp, tt.want)
			}
		})
	}
}

func TestTemperature_StringIn(t *testing.T) {
	type fields struct {
		celsius float64
	}
	type args struct {
		unit Unit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Should print 10°C",
			fields: fields{
				celsius: 10,
			},
			args: args{
				unit: Celsius,
			},
			want: "10°C",
		},
		{
			name: "Should print 50°F",
			fields: fields{
				celsius: 10,
			},
			args: args{
				unit: Fahrenheit,
			},
			want: "50°F",
		},
		{
			name: "Should print empty string for invalid unit",
			fields: fields{
				celsius: 10,
			},
			args: args{
				unit: "Invalid",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temperature := NewFromCelsius(tt.fields.celsius)
			if got := temperature.StringIn(tt.args.unit); got != tt.want {
				t.Errorf("StringIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature_Float64In(t *testing.T) {
	type fields struct {
		celsius float64
	}
	type args struct {
		unit Unit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Should get 10°C",
			fields: fields{
				celsius: 10,
			},
			args: args{
				unit: Celsius,
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "Should get 50°F",
			fields: fields{
				celsius: 10,
			},
			args: args{
				unit: Fahrenheit,
			},
			want:    50,
			wantErr: false,
		},
		{
			name: "Should get empty string for invalid unit",
			fields: fields{
				celsius: 10,
			},
			args: args{
				unit: "Invalid",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temperature := NewFromCelsius(tt.fields.celsius)
			got, err := temperature.Float64In(tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Float64In() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Float64In() got = %v, want %v", got, tt.want)
			}
		})
	}
}
