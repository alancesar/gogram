package numeric

import "testing"

func TestFormat(t *testing.T) {
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
			if got := Format(tt.args.value, tt.args.precision); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
