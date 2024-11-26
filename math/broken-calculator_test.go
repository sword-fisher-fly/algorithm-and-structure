package math

import "testing"

func TestBrokenCalculator(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{x: 68, y: 71},
			want: 34,
		},
		{
			name: "case 2",
			args: args{x: 1, y: 1000000000},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BrokenCalculator(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("BrokenCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
