package volume

import (
	"github.com/alancesar/gogram/measure"
	"reflect"
	"testing"
)

func TestNewFromMilliliter(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
		{
			name: "Should parse from milliliters",
			args: args{
				value: 1000.0,
			},
			want: Volume{
				system:  measure.Metric,
				liters:  1,
				gallons: 0.21996924829908776,
			},
		},
		{
			name: "Should parse from milliliters in one gallon equivalent",
			args: args{
				value: 4546.09,
			},
			want: Volume{
				system:  measure.Metric,
				liters:  4.54609,
				gallons: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromMilliliter(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromMilliliter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromLiter(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
		{
			name: "Should parse from liters",
			args: args{
				value: 10,
			},
			want: Volume{
				system:  measure.Metric,
				liters:  10,
				gallons: 2.1996924829908777,
			},
		},
		{
			name: "Should parse from liters in one gallon equivalent",
			args: args{
				value: 4.54609,
			},
			want: Volume{
				system:  measure.Metric,
				liters:  4.54609,
				gallons: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromLiter(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromLiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromGallon(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
		{
			name: "Should parse from gallons",
			args: args{
				value: 1,
			},
			want: Volume{
				system:  measure.Imperial,
				liters:  4.54609,
				gallons: 1,
			},
		},
		{
			name: "Should parse from gallons in one liter equivalent",
			args: args{
				value: 0.21996924829908777,
			},
			want: Volume{
				system:  measure.Imperial,
				liters:  1,
				gallons: 0.21996924829908777,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromGallon(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromGallon() = %v, want %v", got, tt.want)
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
		want Volume
	}{
		{
			name: "Should parse from '1ml' string",
			args: args{
				input: "1ml",
			},
			want: NewFromMilliliter(1),
		},
		{
			name: "Should parse from '1l' string",
			args: args{
				input: "1l",
			},
			want: NewFromLiter(1),
		},
		{
			name: "Should parse from '1gal' string",
			args: args{
				input: "1gal",
			},
			want: NewFromGallon(1),
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

func TestVolume_IsZero(t *testing.T) {
	type fields struct {
		liters float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Should return true if is empty",
			fields: fields{
				liters: 0,
			},
			want: true,
		},
		{
			name: "Should return false if is not empty",
			fields: fields{
				liters: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewFromLiter(tt.fields.liters)
			if got := v.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_Milliliters(t *testing.T) {
	type fields struct {
		system measure.System
		liters float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should return milliliters properly",
			fields: fields{
				system: measure.Metric,
				liters: 0.001,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Volume{
				system: tt.fields.system,
				liters: tt.fields.liters,
			}
			if got := v.Milliliters(); got != tt.want {
				t.Errorf("Milliliters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_Liters(t *testing.T) {
	type fields struct {
		system measure.System
		liters float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should return liters properly",
			fields: fields{
				system: measure.Metric,
				liters: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Volume{
				system: tt.fields.system,
				liters: tt.fields.liters,
			}
			if got := v.Liters(); got != tt.want {
				t.Errorf("Liters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_Gallons(t *testing.T) {
	type fields struct {
		system  measure.System
		gallons float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should return gallons properly",
			fields: fields{
				system:  measure.Imperial,
				gallons: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Volume{
				system:  tt.fields.system,
				gallons: tt.fields.gallons,
			}
			if got := v.Gallons(); got != tt.want {
				t.Errorf("Gallons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_String(t *testing.T) {
	type fields struct {
		system  measure.System
		liters  float64
		gallons float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should print 1 l",
			fields: fields{
				system: measure.Metric,
				liters: 1,
			},
			want: "1 l",
		},
		{
			name: "Should print 1 ml",
			fields: fields{
				system: measure.Metric,
				liters: 0.001,
			},
			want: "1 ml",
		},
		{
			name: "Should print 1 gal",
			fields: fields{
				system:  measure.Imperial,
				gallons: 1,
			},
			want: "1 gal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Volume{
				system:  tt.fields.system,
				liters:  tt.fields.liters,
				gallons: tt.fields.gallons,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_MarshalJSON(t *testing.T) {
	type fields struct {
		system measure.System
		liters float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Should marshall properly",
			fields: fields{
				system: measure.Metric,
				liters: 100,
			},
			want:    []byte(`"100 l"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Volume{
				system: tt.fields.system,
				liters: tt.fields.liters,
			}
			got, err := m.MarshalJSON()
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

func TestVolume_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Volume
		wantErr bool
	}{
		{
			name: "Should unmarshall properly",
			args: args{
				bytes: []byte("100 l"),
			},
			want:    NewFromLiter(100),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Volume{}
			if err := v.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(*v, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", v, tt.want)
			}
		})
	}
}

func TestVolume_StringIn(t *testing.T) {
	type fields struct {
		liters float64
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
			name: "Should print 1 l",
			fields: fields{
				1,
			},
			args: args{
				unit: LiterUnit,
			},
			want: "1 l",
		},
		{
			name: "Should print 1000 ml",
			fields: fields{
				1,
			},
			args: args{
				unit: MilliliterUnit,
			},
			want: "1000 ml",
		},
		{
			name: "Should print 1 l",
			fields: fields{
				4.54609,
			},
			args: args{
				unit: GallonUnit,
			},
			want: "1 gal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewFromLiter(tt.fields.liters)
			if got := v.StringIn(tt.args.unit); got != tt.want {
				t.Errorf("StringIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_Float64In(t *testing.T) {
	type fields struct {
		liters float64
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
			name: "Should get 1 l",
			fields: fields{
				liters: 1,
			},
			args: args{
				unit: LiterUnit,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Should get 1000 ml",
			fields: fields{
				liters: 1,
			},
			args: args{
				unit: MilliliterUnit,
			},
			want:    1000,
			wantErr: false,
		},
		{
			name: "Should get 1 gal",
			fields: fields{
				liters: 4.54609,
			},
			args: args{
				unit: GallonUnit,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewFromLiter(tt.fields.liters)
			got, err := v.Float64In(tt.args.unit)
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
