package str

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+3-2*4+4",
			args: args{s: "1+3-2*4+4"},
			want: 0,
		},
		{
			name: "-1+3-2*4+4",
			args: args{s: "-1+3-2*4+4"},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.s); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
