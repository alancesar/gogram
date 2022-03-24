package mass

import (
	"github.com/alancesar/gogram/measure"
	"reflect"
	"testing"
)

func TestNewFromMilligram(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse from milligrams",
			args: args{
				value: 1000000.0,
			},
			want: Mass{
				system: measure.Metric,
				grams:  1000.0,
				pounds: 2.2046244201837775,
			},
		},
		{
			name: "Should parse from milligrams to one pound",
			args: args{
				value: 453592,
			},
			want: Mass{
				system: measure.Metric,
				grams:  453.592,
				pounds: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromMilligram(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromMilligram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromGram(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse from grams",
			args: args{
				value: 1000.0,
			},
			want: Mass{
				system: measure.Metric,
				grams:  1000.0,
				pounds: 2.2046244201837775,
			},
		},
		{
			name: "Should parse from grams in one pound equivalent",
			args: args{
				value: 453.592,
			},
			want: Mass{
				system: measure.Metric,
				grams:  453.592,
				pounds: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromGram(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromGram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromKilogram(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse from kilograms",
			args: args{
				value: 1.0,
			},
			want: Mass{
				system: measure.Metric,
				grams:  1000.0,
				pounds: 2.2046244201837775,
			},
		},
		{
			name: "Should parse from kilograms in one pound equivalent",
			args: args{
				value: 0.453592,
			},
			want: Mass{
				system: measure.Metric,
				grams:  453.592,
				pounds: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromKilogram(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromKilogram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromPound(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse from pounds",
			args: args{
				value: 1.0,
			},
			want: Mass{
				system: measure.Imperial,
				grams:  453.592,
				pounds: 1,
			},
		},
		{
			name: "Should parse from pounds in one kilogram",
			args: args{
				value: 2.20462262185,
			},
			want: Mass{
				system: measure.Imperial,
				grams:  999.9991842901852,
				pounds: 2.20462262185,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromPound(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromPound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromOunce(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should parse from ounces",
			args: args{
				value: 1.0,
			},
			want: Mass{
				system: measure.Imperial,
				grams:  28.3495,
				pounds: 0.0625,
			},
		},
		{
			name: "Should parse from ounces in one kilogram",
			args: args{
				value: 35.27396195,
			},
			want: Mass{
				system: measure.Imperial,
				grams:  999.999184301525,
				pounds: 2.204622621875,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromOunce(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromOunce() = %v, want %v", got, tt.want)
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
		want Mass
	}{
		{
			name: "Should parse from '1mg' string",
			args: args{
				input: "1mg",
			},
			want: NewFromMilligram(1),
		},
		{
			name: "Should parse from '1g' string",
			args: args{
				input: "1g",
			},
			want: NewFromGram(1),
		},
		{
			name: "Should parse from '1kg' string",
			args: args{
				input: "1kg",
			},
			want: NewFromKilogram(1),
		},
		{
			name: "Should parse from '1lb' string",
			args: args{
				input: "1lb",
			},
			want: NewFromPound(1),
		},
		{
			name: "Should parse from '1oz' string",
			args: args{
				input: "1oz",
			},
			want: NewFromOunce(1),
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

func TestMass_Milligrams(t *testing.T) {
	type fields struct {
		grams float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should get milligrams properly",
			fields: fields{
				grams: 1,
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewFromGram(tt.fields.grams)
			if got := m.Milligrams(); got != tt.want {
				t.Errorf("Milligrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMass_Grams(t *testing.T) {
	type fields struct {
		grams float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should get grams properly",
			fields: fields{
				grams: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewFromGram(tt.fields.grams)
			if got := m.Grams(); got != tt.want {
				t.Errorf("Grams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMass_Kilograms(t *testing.T) {
	type fields struct {
		grams float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should get grams properly",
			fields: fields{
				grams: 1000,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewFromGram(tt.fields.grams)
			if got := m.Kilograms(); got != tt.want {
				t.Errorf("Kilograms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMass_Pounds(t *testing.T) {
	type fields struct {
		grams float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should get pounds properly",
			fields: fields{
				grams: poundsInGrams,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewFromGram(tt.fields.grams)
			if got := m.Pounds(); got != tt.want {
				t.Errorf("Pounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMass_Ounces(t *testing.T) {
	type fields struct {
		grams float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should get ounces properly",
			fields: fields{
				grams: poundsInGrams / poundsInOunces,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewFromGram(tt.fields.grams)
			if got := m.Ounces(); got != tt.want {
				t.Errorf("Ounces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMass_IsZero(t *testing.T) {
	type fields struct {
		grams float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Should return true if is empty",
			fields: fields{
				grams: 0,
			},
			want: true,
		},
		{
			name: "Should return false if is not empty",
			fields: fields{
				grams: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewFromGram(tt.fields.grams)
			if got := m.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMass_String(t *testing.T) {
	type fields struct {
		system measure.System
		grams  float64
		pounds float64
	}
	tests := []struct {
		name          string
		fields        fields
		wantFormatted string
	}{
		{
			name: "Should print 1.00 g",
			fields: fields{
				system: measure.Metric,
				grams:  1,
			},
			wantFormatted: "1.00 g",
		},
		{
			name: "Should print 1.00 kg",
			fields: fields{
				system: measure.Metric,
				grams:  1000,
			},
			wantFormatted: "1.00 kg",
		},
		{
			name: "Should print 1 mg",
			fields: fields{
				system: measure.Metric,
				grams:  0.001,
			},
			wantFormatted: "1 mg",
		},
		{
			name: "Should print 1 oz",
			fields: fields{
				system: measure.Imperial,
				pounds: 0.0625,
			},
			wantFormatted: "1 oz",
		},
		{
			name: "Should print 1 lb",
			fields: fields{
				system: measure.Imperial,
				pounds: 1,
			},
			wantFormatted: "1.00 lb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mass{
				system: tt.fields.system,
				grams:  tt.fields.grams,
				pounds: tt.fields.pounds,
			}
			if gotFormatted := m.String(); gotFormatted != tt.wantFormatted {
				t.Errorf("String() = %v, want %v", gotFormatted, tt.wantFormatted)
			}
		})
	}
}

func TestMass_MarshalJSON(t *testing.T) {
	type fields struct {
		system measure.System
		grams  float64
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
				grams:  100,
			},
			want:    []byte(`"100.00 g"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mass{
				system: tt.fields.system,
				grams:  tt.fields.grams,
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

func TestMass_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Mass
		wantErr bool
	}{
		{
			name: "Should unmarshall properly",
			args: args{
				bytes: []byte("100 g"),
			},
			want:    NewFromGram(100),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mass{}
			if err := m.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(*m, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", m, tt.want)
			}
		})
	}
}
