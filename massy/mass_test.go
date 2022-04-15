package massy

import (
	"reflect"
	"testing"
)

func TestMass_In(t *testing.T) {
	type fields struct {
		source     Unit
		milligrams int64
	}
	type args struct {
		unit Unit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "Should return equivalent value in milligrams",
			fields: fields{
				source:     Milligram,
				milligrams: 1,
			},
			args: args{
				unit: Milligram,
			},
			want: 1,
		},
		{
			name: "Should return equivalent value in grams",
			fields: fields{
				source:     Gram,
				milligrams: 1000,
			},
			args: args{
				unit: Gram,
			},
			want: 1,
		},
		{
			name: "Should return equivalent value in kilograms",
			fields: fields{
				source:     Kilogram,
				milligrams: 1000000,
			},
			args: args{
				unit: Kilogram,
			},
			want: 1,
		},
		{
			name: "Should return equivalent value in ounces",
			fields: fields{
				source:     Ounce,
				milligrams: 28350,
			},
			args: args{
				unit: Ounce,
			},
			want: 1,
		},
		{
			name: "Should return equivalent value in pounds",
			fields: fields{
				source:     Pound,
				milligrams: 453592,
			},
			args: args{
				unit: Pound,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mass{
				source:     tt.fields.source,
				milligrams: tt.fields.milligrams,
			}
			if got := m.In(tt.args.unit); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		value float64
		unit  Unit
	}
	tests := []struct {
		name string
		args args
		want Mass
	}{
		{
			name: "Should create from milligram",
			args: args{
				value: 1,
				unit:  Milligram,
			},
			want: Mass{
				source:     Milligram,
				milligrams: 1,
			},
		},
		{
			name: "Should create from gram",
			args: args{
				value: 1,
				unit:  Gram,
			},
			want: Mass{
				source:     Gram,
				milligrams: 1000,
			},
		},
		{
			name: "Should create from kilogram",
			args: args{
				value: 1,
				unit:  Kilogram,
			},
			want: Mass{
				source:     Kilogram,
				milligrams: 1000000,
			},
		},
		{
			name: "Should create from ounces",
			args: args{
				value: 1,
				unit:  Ounce,
			},
			want: Mass{
				source:     Ounce,
				milligrams: 28350,
			},
		},
		{
			name: "Should create from pounds",
			args: args{
				value: 1,
				unit:  Pound,
			},
			want: Mass{
				source:     Pound,
				milligrams: 453600,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.value, tt.args.unit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_Symbol(t *testing.T) {
	tests := []struct {
		name string
		u    Unit
		want string
	}{
		{
			name: "Should return mg for Milligram",
			u:    Milligram,
			want: "mg",
		},
		{
			name: "Should return g for Gram",
			u:    Gram,
			want: "g",
		},
		{
			name: "Should return kg for Kilogram",
			u:    Kilogram,
			want: "kg",
		},
		{
			name: "Should return oz for Ounce",
			u:    Ounce,
			want: "oz",
		},
		{
			name: "Should return lb for Pound",
			u:    Pound,
			want: "lb",
		},
		{
			name: "Should return g as default",
			want: "g",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Symbol(); got != tt.want {
				t.Errorf("Symbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_System(t *testing.T) {
	tests := []struct {
		name string
		u    Unit
		want System
	}{
		{
			name: "Should return metric for Milligram",
			u:    Milligram,
			want: Metric,
		},
		{
			name: "Should return metric for Gram",
			u:    Gram,
			want: Metric,
		},
		{
			name: "Should return metric for Kilogram",
			u:    Kilogram,
			want: Metric,
		},
		{
			name: "Should return imperial for Ounce",
			u:    Ounce,
			want: Imperial,
		},
		{
			name: "Should return imperial for Pound",
			u:    Pound,
			want: Imperial,
		},
		{
			name: "Should return metric as default",
			want: Metric,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.System(); got != tt.want {
				t.Errorf("System() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMass_String(t *testing.T) {
	type fields struct {
		source     Unit
		milligrams int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should print 1mg",
			fields: fields{
				source:     Milligram,
				milligrams: 1,
			},
			want: "1mg",
		},
		{
			name: "Should print 1g",
			fields: fields{
				source:     Milligram,
				milligrams: 1000,
			},
			want: "1g",
		},
		{
			name: "Should print 1kg",
			fields: fields{
				source:     Milligram,
				milligrams: 1000000,
			},
			want: "1kg",
		},
		{
			name: "Should print 1oz",
			fields: fields{
				source:     Ounce,
				milligrams: 28350,
			},
			want: "1oz",
		},
		{
			name: "Should print 1lb",
			fields: fields{
				source:     Ounce,
				milligrams: 453600,
			},
			want: "1lb",
		},
		{
			name: "Should print 1.5kg",
			fields: fields{
				source:     Milligram,
				milligrams: 1500000,
			},
			want: "1.5kg",
		},
		{
			name: "Should print 1.59kg",
			fields: fields{
				source:     Milligram,
				milligrams: 1590000,
			},
			want: "1.59kg",
		},
		{
			name: "Should print 1.6kg",
			fields: fields{
				source:     Milligram,
				milligrams: 1599000,
			},
			want: "1.6kg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mass{
				source:     tt.fields.source,
				milligrams: tt.fields.milligrams,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
