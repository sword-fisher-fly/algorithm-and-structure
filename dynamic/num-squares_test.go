package dynamic

import "testing"

func TestNumSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "case 2",
			args: args{n: 12},
			want: 3,
		},
		{
			name: "case 3",
			args: args{n: 13},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSquares(tt.args.n); got != tt.want {
				t.Errorf("NumSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
