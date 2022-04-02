package numeric

import "testing"

func TestFormat(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should print 10.0",
			args: args{
				value: 9.99,
			},
			want: "9.99",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.value); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatWithPrecision(t *testing.T) {
	type args struct {
		value     float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should print 10.0",
			args: args{
				value:     9.99,
				precision: 1,
			},
			want: "10.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatWithPrecision(tt.args.value, tt.args.precision); got != tt.want {
				t.Errorf("FormatWithPrecision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	type args struct {
		input   float64
		decimal int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Should round to 10",
			args: args{
				input:   9.9,
				decimal: 0,
			},
			want: 10,
		},
		{
			name: "Should round to 9.9",
			args: args{
				input:   9.89,
				decimal: 1,
			},
			want: 9.9,
		},
		{
			name: "Should round to 9.99",
			args: args{
				input:   9.989,
				decimal: 2,
			},
			want: 9.99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.input, tt.args.decimal); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
